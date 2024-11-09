package link

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"shortlink/internal/base/errno"
	"shortlink/internal/base/toolkit"
	"strings"
	"time"
)

type FactoryConfig struct {
	Domain            string
	UseSSL            bool
	Whitelist         []string
	MaxAttempts       int // 生成唯一短链接的最大尝试次数
	MaxLinksPerGroup  int
	DefaultFavicon    string
	DefaultGid        string
	DefaultExpiration int
	DefaultCreateType CreateType
	DefaultValidType  ValidType
}

// Validate 这个方法会在配置初始化的时候被调用
func (f FactoryConfig) validate() error {
	var err error

	if f.Domain == "" {
		err = errors.Join(err, errors.New("domain should not be empty"))
	}
	if !toolkit.IsValidDomain(f.Domain) {
		err = errors.Join(err, errors.New("domain should be valid, but is "+f.Domain))
	}
	for _, v := range f.Whitelist {
		if !toolkit.IsValidDomain(v) {
			err = errors.Join(err, errors.New("whitelist domain should be valid, but is "+v))
		}
	}
	if f.DefaultFavicon == "" {
		if !toolkit.IsValidUrl(f.DefaultFavicon) {
			err = errors.Join(err, errors.New("default favicon should be valid url"))
		}
	}
	if f.MaxAttempts < 1 {
		err = errors.Join(
			err,
			errors.New("MaxAttempts should be greater than 1, but is "+fmt.Sprint(f.MaxAttempts)),
		)
	}
	if f.DefaultGid == "" {
		err = errors.Join(err, errors.New("default gid should not be empty"))
	}
	if f.DefaultExpiration < 1 {
		err = errors.Join(
			err,
			errors.New("default expiration should be greater than 1, but is "+fmt.Sprint(f.DefaultExpiration)),
		)
	}
	if f.DefaultCreateType != CreateByConsole && f.DefaultCreateType != CreateByApi {
		err = errors.Join(
			err,
			errors.New("default create type should be 0 or 1, but is "+fmt.Sprint(f.DefaultCreateType)),
		)
	}
	if f.DefaultValidType != ValidTypePermanent && f.DefaultValidType != ValidTypeTemporary {
		err = errors.Join(
			err,
			errors.New("default valid type should be 0 or 1, but is "+fmt.Sprint(f.DefaultValidType)),
		)
	}

	return err
}

// Factory 全局唯一短链接工厂
type Factory struct {
	fc FactoryConfig
}

func NewFactory(fc FactoryConfig) (*Factory, error) {
	if err := fc.validate(); err != nil {
		return &Factory{}, errors.Join(err, errors.New("invalid config passed to link factory"))
	}

	return &Factory{fc: fc}, nil
}

func (f Factory) NewAvailableLink(
	originalUrl string,
	gid string,
	createType *CreateType,
	validType *ValidType,
	startDate *time.Time,
	endDate *time.Time,
	desc string,
	ifExistsFunc func(string) (bool, error),
) (lk *Link, err error) {

	// 白名单校验
	if err = f.verifyWhiteList(originalUrl); err != nil {
		return nil, err
	}

	// 短链接
	var shortUri string
	if shortUri, err = f.genUniqueShortUri(originalUrl, f.fc.MaxAttempts, ifExistsFunc); err != nil {
		return nil, err
	}

	// 完整短链接
	var fullShortUrl string
	if f.fc.UseSSL {
		fullShortUrl = fmt.Sprintf("https://%s/%s", f.fc.Domain, shortUri)
	} else {
		fullShortUrl = fmt.Sprintf("http://%s/%s", f.fc.Domain, shortUri)
	}

	// 标题和图标
	title, favicon, err := toolkit.GetTitleAndFavicon(originalUrl)
	if err == nil {
		// 标题可以为空 但图标有默认值
		if desc == "" {
			desc = title
		}
		if favicon == "" {
			favicon = f.fc.DefaultFavicon
		}
	}

	// 赋予默认值
	if createType == nil {
		createType = &f.fc.DefaultCreateType
	}
	if validType == nil {
		validType = &f.fc.DefaultValidType
	}
	if startDate == nil {
		now := time.Now()
		startDate = &now
	}
	if endDate == nil && *validType == ValidTypeTemporary {
		endDate = &time.Time{}
		*endDate = (*startDate).AddDate(0, 0, f.fc.DefaultExpiration)
	}

	// 有效期
	var validDate *ValidDate
	if validDate, err = NewValidDate(*validType, startDate, endDate); err != nil {
		return nil, err
	}

	return &Link{
		domain:       f.fc.Domain,
		shortUri:     shortUri,
		fullShortUrl: fullShortUrl,
		originalUrl:  originalUrl,
		gid:          gid,
		status:       StatusActive,
		createType:   *createType,
		validDate:    validDate,
		desc:         desc,
		favicon:      favicon,
	}, nil
}

func (f Factory) NewLinkFromDB(
	id uint,
	gid string,
	shortUri string,
	originalUrl string,
	status Status,
	createType CreateType,
	favicon string,
	desc string,
	validDate *ValidDate,
) (*Link, error) {
	// 完整短链接
	var fullShortUrl string
	if f.fc.UseSSL {
		fullShortUrl = fmt.Sprintf("https://%s/%s", f.fc.Domain, shortUri)
	} else {
		fullShortUrl = fmt.Sprintf("http://%s/%s", f.fc.Domain, shortUri)
	}

	return &Link{
		id:           id,
		domain:       f.fc.Domain,
		shortUri:     shortUri,
		fullShortUrl: fullShortUrl,
		originalUrl:  originalUrl,
		gid:          gid,
		status:       status,
		createType:   createType,
		desc:         desc,
		favicon:      favicon,
		validDate:    validDate,
	}, nil

}

func (f Factory) genUniqueShortUri(
	originalUrl string,
	maxAttempts int,
	ifExistsFunc func(string) (bool, error),
) (shortUri string, err error) {
	for i := 0; i < maxAttempts; i++ {
		shortUri = toolkit.HashToBase62(originalUrl + uuid.NewString())
		var exists bool
		if exists, err = ifExistsFunc(shortUri); err != nil {
			return "", err
		}
		if !exists {
			return shortUri, nil
		}
	}
	return "", errno.LinkTooManyAttempts
}

func (f Factory) verifyWhiteList(originUrl string) error {
	whitelist := f.fc.Whitelist
	if whitelist == nil || len(whitelist) == 0 {
		return nil
	}

	domain := toolkit.ExtractDomain(originUrl)
	if domain == "" {
		return errno.LinkInvalidOriginalUrl
	}
	for _, v := range whitelist {
		// 支持子域名
		if domain == v || strings.HasSuffix(domain, "."+v) {
			return nil
		}
	}
	return errno.LinkDisallowedDomain
}

func (f Factory) CheckGroupLinkCount(count int) error {
	if count >= f.fc.MaxLinksPerGroup {
		return errno.LinkGroupLinkCountExceed
	}
	return nil
}
