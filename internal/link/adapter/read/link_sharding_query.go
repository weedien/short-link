package read

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"shortlink/internal/base/errno"
	"shortlink/internal/base/types"
	"shortlink/internal/link/adapter/po"
	"shortlink/internal/link/app/query"
	"shortlink/internal/link/domain/link"
	"time"
)

type LinkShardingQuery struct {
	linkFactory *link.Factory
	db          *gorm.DB
}

func NewLinkShardingQuery(db *gorm.DB, factory *link.Factory) LinkShardingQuery {
	return LinkShardingQuery{
		linkFactory: factory,
		db:          db,
	}
}

func (q LinkShardingQuery) PageLink(ctx context.Context, param query.PageLink) (res *types.PageResp[query.Link], err error) {
	rawSql := `
SELECT t.*, COALESCE(s.today_pv, 0) AS todayPv, COALESCE(s.today_uv, 0) AS todayUv, COALESCE(s.today_uip, 0) AS todayUip
FROM t_link t
LEFT JOIN t_link_stats_today s ON t.short_uri = s.short_uri AND s.date = current_date
WHERE t.delete_time is null
`

	var args []interface{}
	if param.Gid != nil {
		rawSql += " AND t.gid = ?"
		args = append(args, *param.Gid)
	}
	if param.OrderTag != nil {
		rawSql += `
ORDER BY 
    CASE 
        WHEN ? = 'todayPv' THEN todayPv
        WHEN ? = 'todayUv' THEN todayUv
        WHEN ? = 'todayUip' THEN todayUip
        WHEN ? = 'totalPv' THEN t.total_pv
        WHEN ? = 'totalUv' THEN t.total_uv
        WHEN ? = 'totalUip' THEN t.total_uip
        ELSE t.create_time
    END DESC
`
		orderTag := *param.OrderTag
		args = append(args, orderTag, orderTag, orderTag, orderTag, orderTag, orderTag)
	} else {
		rawSql += " ORDER BY t.create_time DESC"
	}

	rawSql += " LIMIT ? OFFSET ?"
	args = append(args, param.Limit(), param.Offset())

	var records []query.Link
	if err = q.db.WithContext(ctx).Raw(rawSql, args...).Scan(&records).Error; err != nil {
		return
	}

	var total int64
	qu := q.db.WithContext(ctx).Model(&po.Link{})
	if param.Gid != nil {
		qu = qu.Where("gid = ?", *param.Gid)
	}

	if err = qu.Count(&total).Error; err != nil {
		return
	}

	res = &types.PageResp[query.Link]{
		Current: param.Current,
		Size:    param.Size,
		Total:   total,
		Records: records,
	}
	return
}

func (q LinkShardingQuery) ListGroupLinkCount(ctx context.Context, gids []string) (res []query.GroupLinkCount, err error) {
	if err = q.db.WithContext(ctx).
		Table("link").Select("gid, COUNT(*) AS count").Where("gid IN ?", gids).Group("gid").
		Find(&res).Error; err != nil {
		return
	}
	return
}

// PageRecycleBin 分页查询回收站中的短链接
func (q LinkShardingQuery) PageRecycleBin(
	ctx context.Context,
	param query.PageRecycleBin,
) (res *types.PageResp[query.Link], err error) {
	rawSql := `
SELECT t.*, COALESCE(s.today_pv, 0) AS todayPv, COALESCE(s.today_uv, 0) AS todayUv, COALESCE(s.today_uip, 0) AS todayUip
FROM t_link t
LEFT JOIN t_link_stats_today s ON t.short_uri = s.short_uri AND s.date = current_date
WHERE t.gid IN (?) AND t.status = 1
ORDER BY t.update_time
LIMIT ? OFFSET ?;
`
	records := make([]query.Link, 0)
	if err = q.db.WithContext(ctx).Raw(rawSql, param.Gids, param.Limit(), param.Offset()).Scan(&records).Error; err != nil {
		return
	}

	var total int64
	if err = q.db.WithContext(ctx).Model(&po.Link{}).Where("gid IN (?)", param.Gids).Count(&total).Error; err != nil {
		return
	}

	res = &types.PageResp[query.Link]{
		Current: param.Current,
		Size:    param.Size,
		Total:   total,
		Records: records,
	}

	return
}

func (q LinkShardingQuery) GetLink(ctx context.Context, shortUri string) (lk *link.Link, err error) {
	linkGotoPo := po.LinkGoto{}
	if err = q.db.WithContext(ctx).Where("short_uri = ?", shortUri).First(&linkGotoPo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errno.LinkNotExists
		}
		return
	}

	linkPo := po.Link{}
	if err = q.db.WithContext(ctx).Where("short_uri = ? AND gid = ?", shortUri, linkGotoPo.Gid).First(&linkPo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errno.LinkNotExists
		}
		return
	}

	// 将持久化对象转换为领域模型
	lk = &link.Link{}

	var start *time.Time
	var end *time.Time
	if linkPo.StartDate.Valid {
		*start = linkPo.StartDate.Time
	}
	if linkPo.EndDate.Valid {
		*end = linkPo.EndDate.Time
	}

	validDate := &link.ValidDate{}
	if validDate, err = link.NewValidDate(linkPo.ValidType, start, end); err != nil {
		return
	}

	if lk, err = q.linkFactory.NewLinkFromDB(linkPo.ID, linkPo.Gid, linkPo.ShortUri, linkPo.OriginalUrl, linkPo.Status,
		linkPo.CreateType, linkPo.Favicon, linkPo.Desc, validDate); err != nil {
		return
	}

	return
}

// GetOriginalUrlByShortUri 根据短链接获取原始链接
//func (q LinkShardingQuery) GetOriginalUrlByShortUri(ctx context.Context, shortUri string) (int, *string, error) {
//
//	getFromCache := func(key string) (string, error) {
//		val, err := q.rdb.Get(ctx, key).Result()
//		if errors.Is(err, redis.Nil) {
//			return "", nil
//		}
//		return val, err
//	}
//
//	// 尝试从缓存中获取短链的原始链接，存在则直接返回
//	key := fmt.Sprintf(constant.GotoLinkKey, shortUri)
//	originalUrl, err := getFromCache(key)
//	if err != nil {
//		return "", err
//	}
//	if originalUrl != "" {
//		return originalUrl, nil
//	}
//
//	// 如果缓存中没有，判断是否存在于布隆过滤器中，不存在返回 not found
//	exists, err := q.rdb.BFExists(ctx, cache.ShortUriCreateBloomFilter, shortUri).Result()
//	if err != nil {
//		return "", err
//	}
//	if !exists {
//		return "", error_no.LinkNotExists
//	}
//
//	// 从缓存中获取 GotoIsNullLink 的值，如果存在则意味着短链接失效，返回 not found
//	key = fmt.Sprintf(constant.GotoIsNullLinkKey, shortUri)
//	gotoIsNullLink, err := getFromCache(key)
//	if err != nil {
//		return "", err
//	}
//	if gotoIsNullLink != "" {
//		return "", error_no.LinkExpired
//	}
//
//	// 获取分布式锁，并进行二次锁判定，尝试从数据库中获取原始链接，并写入缓存
//	lockKey := fmt.Sprintf(constant.LockGotoLinkKey, shortUri)
//	_, err = q.locker.Acquire(ctx, lockKey, constant.DefaultTimeOut)
//	if errors.Is(err, redislock.ErrNotObtained) {
//		return "", error_no.LockAcquireFailed
//	}
//	defer func() {
//		if err = q.locker.Release(ctx, lockKey); err != nil {
//			err = error_no.LockReleaseFailed
//		}
//	}()
//
//	// ------------- 在查询数据库之前再进行一次判断，防止缓存击穿 -------------
//	// 在高并发场景下，会存在多个线程同时竞争锁，但只有一个线程能够获取锁
//	// 第一个线程执行结束后，其他线程再次尝试获取锁，此时缓存中已经有值，直接返回
//	// TODO: 优化方案，可以使用分布式锁的续租功能，避免锁过期导致的缓存击穿
//	// TODO: 目前我是以Java的思维来写的，Go的锁机制可能有更好的解决方案
//	key = fmt.Sprintf(constant.GotoLinkKey, shortUri)
//	originalUrl = q.rdb.Get(ctx, key).String()
//	if originalUrl != "" {
//		return originalUrl, nil
//	}
//
//	key = fmt.Sprintf(constant.GotoIsNullLinkKey, shortUri)
//	gotoIsNullLink = q.rdb.Get(ctx, key).String()
//	if gotoIsNullLink != "" {
//		return "", error_no.LinkExpired
//	}
//
//	// 当短链接存在且有效时 返回 true, originalUrl, nil
//	// 当短链接处于失效/停用/回收站状态时 返回 false, nil, nil
//	fetchFromDB := func() (interface{}, error) {
//		linkGotoPo, err := q.linkDao.GetLinkGoto(ctx, shortUri)
//		if err != nil {
//			return nil, err
//		}
//		if linkGotoPo == nil {
//			return nil, nil
//		}
//		linkPo, err := q.linkDao.GetLink(ctx, link.Identifier{ShortUri: shortUri, Gid: linkGotoPo.Gid})
//		if err != nil {
//			return nil, err
//		}
//		return linkPo, nil
//	}
//
//	// 查询数据库
//	linkGotoPo, err := q.linkDao.GetLinkGoto(ctx, shortUri)
//	if err != nil {
//		return "", err
//	}
//	// 数据库中不存在这个短链接
//	if linkGotoPo == nil {
//		key = fmt.Sprintf(constant.GotoIsNullLinkKey, shortUri)
//		q.rdb.SetEx(ctx, key, "-", constant.DefaultExpiration)
//		return "", error_no.LinkNotExists
//	}
//	// 数据库中存在短链接，则查询link表获取原始链接
//	linkPo, err := q.linkDao.GetLink(ctx, link.Identifier{ShortUri: shortUri, Gid: linkGotoPo.Gid})
//	if err != nil {
//		return "", err
//	}
//	// 短链接为空或者已经过期
//	if linkPo == nil || linkPo.EndTime.Before(time.Now()) {
//		// 写入缓存
//		key = fmt.Sprintf(constant.GotoIsNullLinkKey, shortUri)
//		q.rdb.SetEx(ctx, key, "-", constant.DefaultExpiration)
//		return "", error_no.LinkExpired
//	}
//	// 查询到有效的原始链接，写入缓存
//	key = fmt.Sprintf(constant.GotoLinkKey, shortUri)
//	err = q.rdb.SetEx(ctx, key, linkPo.OriginalUrl, toolkit.GetLinkCacheExpiration(linkPo.EndTime)).Err()
//	if err != nil {
//		return "", err
//	}
//	return linkPo.OriginalUrl, err
//}

// GetWithCacheAndLock 一个通用的方法，先从缓存中查找，如果不存在，使用分布式锁，进行双重锁判定，从数据库查，然后写入缓存
//
// 能从缓存中查到的短链一定是有效的吗？是的
//func GetWithCacheAndLock[T any](
//	ctx context.Context,
//	cacheKey string,
//	lockKey string,
//	exceptKey string, // 用于从 bloom filter 中排除已失效的数据
//	expiration time.Duration,
//	bloomFilter string,
//	fetchFromDB func() (T, error),
//) (val T, err error) {
//	// step1: 尝试从缓存中取值
//	cacheVal, err := q.rdb.Get(ctx, cacheKey).Result()
//	if err == nil && cacheVal != "" {
//		// Assuming the cached value is a JSON string, unmarshal it into the generic type T
//		err = sonic.Unmarshal([]byte(cacheVal), &val)
//		if err == nil {
//			return val, nil
//		}
//	}
//
//	// step2: 缓存中没有，通过 布隆过滤器+失效缓存 判断是否存在
//	// case1: bloom filter 中不存在 则数据一定不存在
//	exists := false
//	if exists, err = q.rdb.BFExists(ctx, bloomFilter, cacheKey).Result(); err != nil {
//		return val, err
//	}
//	if !exists {
//		return val, error_no.LinkNotExists
//	}
//	// case2: bloom filter 中存在 但失效缓存中也存在 则数据已失效
//	if cacheVal, err = q.rdb.Get(ctx, exceptKey).Result(); err != nil {
//		if errors.Is(err, redis.Nil) {
//			// 数据存在且有效
//		} else {
//			// redis 异常
//			return val, errors.Join(err, error_no.RedisError)
//		}
//	} else {
//		// 数据已失效
//		return val, error_no.LinkExpired
//	}
//
//	// step3: 获取分布式锁
//	acquired := false
//	if acquired, err = q.locker.Acquire(ctx, lockKey, expiration); err != nil {
//		return val, err
//	}
//	if !acquired {
//		return val, error_no.LockAcquireFailed
//	}
//	defer func(locker lock.DistributedLock, ctx context.Context, key string) {
//		if releaseErr := locker.Release(ctx, key); releaseErr != nil {
//			err = releaseErr
//		}
//	}(q.locker, ctx, lockKey)
//
//	// 双重判断，防止缓存击穿
//	if cacheVal, err = q.rdb.Get(ctx, cacheKey).Result(); err != nil {
//		if errors.Is(err, redis.Nil) {
//			return val, error_no.LinkNotExists
//		} else {
//			return val, errors.Join(err, error_no.RedisError)
//		}
//	} else if cacheVal != "" {
//		err = sonic.Unmarshal([]byte(cacheVal), &val)
//		if err == nil {
//			return val, nil
//		}
//	}
//	if cacheVal, err = q.rdb.Get(ctx, exceptKey).Result(); err != nil {
//		if errors.Is(err, redis.Nil) {
//			// 数据存在且有效
//		} else {
//			// redis 异常
//			return val, errors.Join(err, error_no.RedisError)
//		}
//	} else {
//		// 数据已失效
//		return val, error_no.LinkExpired
//	}
//
//	// 从数据库中获取
//	var res T
//	if res, err = fetchFromDB(); err != nil {
//		return val, err
//	}
//	if res == nil {
//		// 数据被删除 硬删除/软删除
//		return val, nil
//	}
//
//	// 写入缓存
//	var jsonBytes []byte
//	jsonBytes, err = sonic.Marshal(res)
//	if err != nil {
//		return val, err
//	}
//	err = q.rdb.SetEx(ctx, cacheKey, string(jsonBytes), expiration).Err()
//	if err != nil {
//		return val, err
//	}
//
//	return res, nil
//}
