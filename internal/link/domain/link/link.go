package link

import (
	"errors"
	"shortlink/internal/base/errno"
	"time"
)

// Identifier 短链接唯一标识
//
// 实际上 shortUri 可以确定唯一短链接，而 gid 在分库分表场景下作为分片键，
// 所以 gid 的意义在于避免扩散查询
type Identifier struct {
	// 分组ID
	Gid string
	// 完整短链接
	ShortUri string
}

// Link 短链接
// 由 shortUri 确定唯一短链接
type Link struct {
	id           uint
	domain       string
	shortUri     string
	fullShortUrl string
	originalUrl  string
	gid          string
	status       Status
	createType   CreateType
	desc         string
	favicon      string
	validDate    *ValidDate
}

func (lk Link) ID() uint {
	return lk.id
}

func (lk Link) CreateType() CreateType {
	return lk.createType
}

func (lk Link) Gid() string {
	return lk.gid
}

func (lk Link) Favicon() string {
	return lk.favicon
}

func (lk Link) OriginalUrl() string {
	return lk.originalUrl
}

func (lk Link) ShortUri() string {
	return lk.shortUri
}

func (lk Link) Status() Status {
	return lk.status
}

func (lk Link) RecoverFromRecycleBin() {
	lk.status = StatusActive
}
func (lk Link) SaveToRecycleBin() {
	lk.status = StatusDisabled
}

func (lk Link) FullShortUrl() string {
	return lk.fullShortUrl
}

func (lk Link) ValidDate() *ValidDate {
	return lk.validDate
}

func (lk Link) Desc() string {
	return lk.desc
}

// Update 更新短链接信息
func (lk Link) Update(
	gid string,
	originalUrl string,
	status Status,
	validType *ValidType,
	validEndDate *time.Time,
	desc *string,
) error {
	if gid != "" {
		lk.gid = gid
	}
	if originalUrl != "" {
		lk.originalUrl = originalUrl
	}
	if status != "" {
		lk.status = status
	}
	if validType != nil {
		if *validType != ValidTypePermanent && *validType != ValidTypeTemporary {
			return errors.New("invalid validType")
		}
		lk.validDate.validType = *validType
	}
	if validEndDate != nil {
		if validEndDate.Before(time.Now()) {
			return errors.New("endDate should be after startDate")
		}
		lk.validDate.endDate = validEndDate
	}
	if desc != nil {
		lk.desc = *desc
	}
	return nil
}

type CacheValue struct {
	OriginalUrl string     `json:"originalUrl"`
	NeverExpire bool       `json:"neverExpire"`
	StartTime   *time.Time `json:"startTime"`
	EndTime     *time.Time `json:"endTime"`
	Status      Status     `json:"status"`
}

func NewCacheValue(lk *Link) *CacheValue {
	return &CacheValue{
		OriginalUrl: lk.OriginalUrl(),
		NeverExpire: lk.ValidDate().NeverExpire(),
		StartTime:   lk.ValidDate().StartTime(),
		EndTime:     lk.ValidDate().EndTime(),
		Status:      lk.Status(),
	}
}

func (c CacheValue) Validate() (bool, error) {
	if c.Status == StatusActive {
		if c.NeverExpire {
			return true, nil
		}
		if c.StartTime.Before(time.Now()) && c.EndTime.After(time.Now()) {
			return true, nil
		}
		return false, errno.LinkExpired
	}
	switch {
	case c.Status == StatusReserved:
		return false, errno.LinkReserved
	case c.Status == StatusForbidden:
		return false, errno.LinkForbidden
	case c.Status == StatusDisabled:
		return false, errno.LinkDisabled
	case c.Status == StatusExpired:
		return false, errno.LinkExpired
	}
	return false, nil
}

func (c CacheValue) Expiration() time.Duration {
	if c.NeverExpire {
		return 0
	}
	return c.EndTime.Sub(time.Now())
}
