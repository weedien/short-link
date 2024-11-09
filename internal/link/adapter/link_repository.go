package adapter

import (
	"context"
	"database/sql"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"reflect"
	"shortlink/internal/base/cache"
	"shortlink/internal/base/errno"
	"shortlink/internal/link/adapter/assembler"
	"shortlink/internal/link/adapter/po"
	"shortlink/internal/link/common/constant"
	"shortlink/internal/link/domain/link"
	"strconv"
	"time"
)

type LinkRepository struct {
	db               *gorm.DB
	distributedCache cache.DistributedCache
	assembler        assembler.LinkAssembler
}

func NewLinkRepository(
	linkFactory *link.Factory,
	db *gorm.DB,
	distributedCache cache.DistributedCache,
) *LinkRepository {
	return &LinkRepository{
		db:               db,
		distributedCache: distributedCache,
		assembler:        assembler.NewLinkAssembler(linkFactory),
	}
}

func (r LinkRepository) ShortUriExists(ctx context.Context, shortUri string) (bool, error) {
	return r.distributedCache.ExistsInBloomFilter(
		ctx,
		cache.ShortUriCreateBloomFilter,
		shortUri,
		constant.GotoIsNullLinkKey+shortUri,
	)
}

func (r LinkRepository) CountLinksByGid(ctx context.Context, gid string) (int, error) {
	cnt, err := r.distributedCache.HGet(ctx, constant.LinkGroupCountKey, gid)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// 查询数据库
			var count int64
			if err = r.db.WithContext(ctx).Model(&po.Link{}).
				Where("gid = ?", gid).Count(&count).Error; err != nil {
				return 0, err
			}
			return int(count), nil
		} else {
			return 0, err
		}
	}

	if cnt == "" {
		return 0, nil
	} else {
		return strconv.Atoi(cnt)
	}
}

func (r LinkRepository) CreateLink(ctx context.Context, lk *link.Link) error {
	// 大概消耗时间：50ms
	//if exists, err := r.distributedCache.ExistsInBloomFilter(
	//	ctx,
	//	cache.ShortUriCreateBloomFilter,
	//	lk.ShortUri(),
	//	constant.GotoIsNullLinkKey+lk.ShortUri(),
	//); err != nil {
	//	return err
	//} else if exists {
	//	return errno.LinkAlreadyExists
	//}

	linkPo := r.assembler.LinkEntityToLinkPo(lk)
	if err := r.db.WithContext(ctx).Create(&linkPo).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errno.LinkAlreadyExists
		}
		return err
	}

	// 放入缓存和布隆过滤器
	cacheValue := link.NewCacheValue(lk)
	if err := r.distributedCache.SafePut(
		ctx,
		constant.GotoLinkKey+lk.ShortUri(),
		cacheValue,
		cacheValue.Expiration(),
		cache.ShortUriCreateBloomFilter,
		lk.ShortUri(),
	); err != nil {
		return err
	}

	// 统计分组数量 hash
	username := ctx.Value("username").(string)
	if _, err := r.distributedCache.HIncrBy(
		ctx,
		constant.LinkGroupCountKey+username,
		lk.Gid(),
		1,
	); err != nil {
		return err
	}

	return nil
}

func (r LinkRepository) CreateLinkBatch(ctx context.Context, links []*link.Link) error {

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, lk := range links {
			linkPo := r.assembler.LinkEntityToLinkPo(lk)
			if err := tx.Create(&linkPo).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	for _, lk := range links {
		err := r.distributedCache.SafePut(
			ctx,
			constant.GotoLinkKey+lk.ShortUri(),
			link.NewCacheValue(lk),
			lk.ValidDate().Expiration(),
			cache.ShortUriCreateBloomFilter,
			lk.ShortUri(),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r LinkRepository) UpdateLink(
	ctx context.Context,
	shortUri string,
	updateFn func(ctx context.Context, link *link.Link) (*link.Link, error),
) error {
	var linkPo po.Link
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := r.db.WithContext(ctx).
			Model(&linkPo).
			Where("short_uri = ?", shortUri).
			First(&linkPo).Error; err != nil {
			return err
		}

		lk := r.assembler.LinkPoToLinkEntity(linkPo)

		var err error
		if lk, err = updateFn(ctx, lk); err != nil {
			return err
		}

		linkPo = r.assembler.LinkEntityToLinkPo(lk)
		return r.db.WithContext(ctx).Updates(&linkPo).Error

	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.LinkNotExists
		}
	}
	return err
}

func (r LinkRepository) SaveToRecycleBin(ctx context.Context, id link.Identifier) error {
	// 持久化操作
	var linkPo po.Link
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 查询
		if err := r.db.WithContext(ctx).
			Model(&linkPo).
			Where("short_uri = ?", id.ShortUri).
			First(&linkPo).Error; err != nil {
			return err
		}

		// 修改
		linkPo.RecycleTime = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}

		// 更新
		if err := r.db.WithContext(ctx).Save(&linkPo).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.LinkNotExists
		}
		return err
	}

	// 修改缓存中的状态为已删除
	cacheKey := constant.GotoLinkKey + id.ShortUri
	if err = r.modifyCacheValueStatus(ctx, cacheKey, link.StatusDeleted); err != nil {
		return err
	}
	return nil
}

func (r LinkRepository) RemoveFromRecycleBin(ctx context.Context, id link.Identifier) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		var linkPo po.Link
		if err = r.db.WithContext(ctx).
			Model(&linkPo).
			Where("short_uri = ?", id.ShortUri).
			First(&linkPo).Error; err != nil {
			return
		}

		//如果短链接状态不是回收站状态，返回错误
		if !linkPo.RecycleTime.Valid {
			return errno.LinkInvalidStatus
		}

		return r.db.Delete(&linkPo).Error
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.LinkNotExists
		}
		return err
	}

	// 删除缓存
	if err = r.distributedCache.SafeDelete(
		ctx, constant.GotoLinkKey+id.ShortUri,
		constant.GotoIsNullLinkKey+id.ShortUri,
	); err != nil {
		return err
	}

	return nil
}

func (r LinkRepository) RecoverFromRecycleBin(ctx context.Context, id link.Identifier) error {
	var linkPo po.Link
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := r.db.WithContext(ctx).Model(&linkPo).
			Where("short_uri = ?", id.ShortUri).
			First(&linkPo).Error; err != nil {
			return err
		}

		// 如果短链接状态不是回收站状态，返回错误
		if !linkPo.RecycleTime.Valid {
			return errno.LinkInvalidStatus
		}

		// 修改
		linkPo.RecycleTime = sql.NullTime{
			Valid: false,
		}

		return r.db.Save(&linkPo).Error
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.LinkNotExists
		}
		return err
	}

	cacheKey := constant.GotoLinkKey + id.ShortUri
	if err = r.modifyCacheValueStatus(ctx, cacheKey, link.Status(linkPo.Status)); err != nil {
		return err
	}
	return nil
}

func (r LinkRepository) modifyCacheValueStatus(ctx context.Context, cacheKey string, status link.Status) (err error) {
	var res interface{}
	cacheValue := &link.CacheValue{}
	if res, err = r.distributedCache.Get(ctx, cacheKey, reflect.TypeOf(link.CacheValue{})); err != nil {
		return err
	}
	cacheValue = res.(*link.CacheValue)
	cacheValue.Status = status
	if err = r.distributedCache.Put(ctx, cacheKey, cacheValue, cacheValue.Expiration()); err != nil {
		return err
	}
	return
}
