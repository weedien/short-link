package cache

//func GetWithCacheAndLock[T any](
//	ctx context.Context,
//	locker lock.DistributedLock,
//	cacheKey string,
//	lockKey string,
//	exceptKey string, // 用于从 bloom filter 中排除已失效的数据
//	expiration time.Duration,
//	bloomFilter string,
//	fetchFn func() (T, error),
//) (val T, err error) {
//	// step1: 尝试从缓存中取值
//	cacheVal, err := rdb.Get(ctx, cacheKey).Result()
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
//	if exists, err = rdb.BFExists(ctx, bloomFilter, cacheKey).Result(); err != nil {
//		return val, err
//	}
//	if !exists {
//		return val, errno.RedisKeyNotExist
//	}
//	// case2: bloom filter 中存在 但失效缓存中也存在 则数据已失效
//	if cacheVal, err = rdb.Get(ctx, exceptKey).Result(); err != nil {
//		if errors.Is(err, redis.Nil) {
//			// 数据存在且有效
//		} else {
//			// redis 异常
//			return val, errors.Join(err, errno.RedisError)
//		}
//	} else {
//		// 数据已失效
//		return val, errno.RedisKeyNotExist
//	}
//
//	// step3: 获取分布式锁
//	acquired := false
//	if acquired, err = locker.Acquire(ctx, lockKey, expiration); err != nil {
//		return val, err
//	}
//	if !acquired {
//		return val, errno.LockAcquireFailed
//	}
//	defer func(locker lock.DistributedLock, ctx context.Context, key string) {
//		if releaseErr := locker.Release(ctx, key); releaseErr != nil {
//			err = releaseErr
//		}
//	}(locker, ctx, lockKey)
//
//	// 双重判断，防止缓存击穿
//	if cacheVal, err = rdb.Get(ctx, cacheKey).Result(); err != nil {
//		if errors.Is(err, redis.Nil) {
//			return val, errno.RedisKeyNotExist
//		} else {
//			return val, errors.Join(err, errno.RedisError)
//		}
//	} else if cacheVal != "" {
//		err = sonic.Unmarshal([]byte(cacheVal), &val)
//		if err == nil {
//			return val, nil
//		}
//	}
//	if cacheVal, err = rdb.Get(ctx, exceptKey).Result(); err != nil {
//		if errors.Is(err, redis.Nil) {
//			// 数据存在且有效
//		} else {
//			// redis 异常
//			return val, errors.Join(err, errno.RedisError)
//		}
//	} else {
//		// 数据已失效
//		return val, errno.RedisKeyNotExist
//	}
//
//	// 从数据库中获取
//	var res T
//	if res, err = fetchFn(); err != nil {
//		return val, err
//	}
//	isNil := false
//	if isNil, err = isNilOrEmpty(res); isNil || err != nil {
//		// 数据被删除 硬删除/软删除
//		return val, err
//	}
//
//	// 写入缓存
//	var jsonBytes []byte
//	jsonBytes, err = sonic.Marshal(res)
//	if err != nil {
//		return val, err
//	}
//	err = rdb.SetEx(ctx, cacheKey, string(jsonBytes), expiration).Err()
//	if err != nil {
//		return val, err
//	}
//
//	return res, nil
//}
