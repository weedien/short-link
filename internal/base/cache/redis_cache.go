package cache

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"reflect"
	"shortlink/internal/base"
	"shortlink/internal/base/errno"
	"shortlink/internal/base/lock"
	"time"
)

type RedisDistributedCache struct {
	rdb    *redis.Client
	locker lock.DistributedLock
}

func NewRedisDistributedCache(rdb *redis.Client, locker lock.DistributedLock) *RedisDistributedCache {
	if rdb == nil {
		panic("nil rdb")
	}
	if locker == nil {
		panic("nil locker")
	}
	return &RedisDistributedCache{rdb: rdb, locker: locker}
}

func (r RedisDistributedCache) Get(ctx context.Context, key string, valueType reflect.Type) (interface{}, error) {
	value, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errno.RedisKeyNotExist
		}
		return nil, err
	}

	if valueType == reflect.TypeOf("") {
		return value, nil
	}

	result := reflect.New(valueType).Interface()
	if err = sonic.Unmarshal([]byte(value), result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r RedisDistributedCache) HGet(ctx context.Context, key string, gid string) (string, error) {
	return r.rdb.HGet(ctx, key, gid).Result()
}

func (r RedisDistributedCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.rdb.HGetAll(ctx, key).Result()
}

func (r RedisDistributedCache) Put(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	valueBytes, err := sonic.Marshal(value)
	if err != nil {
		return err
	}
	return r.rdb.Set(ctx, key, string(valueBytes), expiration).Err()
}

func (r RedisDistributedCache) PutIfAbsent(ctx context.Context, key string, value interface{}) (bool, error) {
	luaScript := `
		if redis.call("EXISTS", KEYS[1]) == 0 then
			redis.call("SET", KEYS[1], ARGV[1])
			return 1
		else
			return 0
		end
	`

	valueBytes, err := sonic.Marshal(value)
	if err != nil {
		return false, err
	}

	result, err := r.rdb.Eval(ctx, luaScript, []string{key}, valueBytes).Result()
	if err != nil {
		return false, err
	}

	if result.(int) == 1 {
		return true, nil
	}
	return false, nil
}

func (r RedisDistributedCache) Delete(ctx context.Context, key string) (bool, error) {
	result, err := r.rdb.Del(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

func (r RedisDistributedCache) DeleteMultiple(ctx context.Context, keys []string) (int, error) {
	result, err := r.rdb.Del(ctx, keys...).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

func (r RedisDistributedCache) HasKey(ctx context.Context, key string) (bool, error) {
	result, err := r.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}
func (r RedisDistributedCache) GetInstance() interface{} {
	return r.rdb
}

func (r RedisDistributedCache) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	return r.rdb.HIncrBy(ctx, key, field, incr).Result()
}

func (r RedisDistributedCache) SafeGet(
	ctx context.Context,
	key string,
	valueType reflect.Type,
	cacheLoader Loader,
	expiration time.Duration,
) (interface{}, error) {
	return r.SafeGetWithCacheGetIfAbsent(ctx, key, valueType, cacheLoader, expiration, "", "", "", nil)
}

func (r RedisDistributedCache) SafeGetWithBloomFilter(
	ctx context.Context,
	key string,
	valueType reflect.Type,
	cacheLoader Loader,
	expiration time.Duration,
	bloomFilter string,
	bloomKey string,
) (interface{}, error) {
	return r.SafeGetWithCacheGetIfAbsent(ctx, key, valueType, cacheLoader, expiration, bloomFilter, bloomKey, "", nil)
}

func (r RedisDistributedCache) SafeGetWithCacheCheckFilter(
	ctx context.Context,
	key string,
	valueType reflect.Type,
	cacheLoader Loader,
	expiration time.Duration,
	bloomFilter string,
	bloomKey string,
	exceptBloomKey string,
) (interface{}, error) {
	return r.SafeGetWithCacheGetIfAbsent(ctx, key, valueType, cacheLoader, expiration, bloomFilter, bloomKey, exceptBloomKey, nil)
}

func (r RedisDistributedCache) SafeGetWithCacheGetIfAbsent(
	ctx context.Context,
	key string,
	valueType reflect.Type,
	cacheLoader Loader,
	expiration time.Duration,
	bloomFilter string,
	bloomKey string,
	exceptBloomKey string,
	cacheGetIfAbsent GetIfAbsent,
) (interface{}, error) {
	// step1 从缓存中取值
	result, err := r.Get(ctx, key, valueType)
	if err != nil {
		if errors.Is(err, errno.RedisKeyNotExist) {

		} else {
			return nil, err
		}
	}
	isNil := false
	if isNil, err = isNilOrEmpty(result); err != nil {
		return nil, err
	}
	if !isNil {
		return result, nil
	}

	// step2 缓存中为空 可能是因为缓存失效 判断是否存在于布隆过滤器中

	// -1: 布隆过滤器不存在 需要查询数据库
	// 0: key 不在布隆过滤器中 不需要查询数据库
	// 1: key 在布隆过滤器中 需要查询数据库
	checkBloom := 0
	if checkBloom, err = r.CheckBloomFilter(ctx, bloomFilter, bloomKey, exceptBloomKey); err != nil {
		return nil, err
	}
	if checkBloom == 0 {
		return nil, errno.RedisKeyNotExist
	}

	// step3 获取分布式锁
	acquired := false
	lockKey := r.defaultLockKey(key)
	if acquired, err = r.locker.Acquire(ctx, lockKey, DefaultTimeOut); err != nil {
		return result, err
	}
	if !acquired {
		return result, errno.LockAcquireFailed
	}
	defer func(locker lock.DistributedLock, ctx context.Context, key string) {
		if releaseErr := locker.Release(ctx, key); releaseErr != nil {
			err = releaseErr
		}
	}(r.locker, ctx, lockKey)

	// 双重判断，防止缓存击穿
	if result, err = r.Get(ctx, key, valueType); err != nil {
		if errors.Is(err, errno.RedisKeyNotExist) {

		} else {
			return nil, err
		}
	}
	if isNil, err = isNilOrEmpty(result); err != nil {
		return nil, err
	}
	if isNil {
		// 从数据库中获取
		if result, err = r.loadAndSet(ctx, key, cacheLoader, expiration, bloomFilter, bloomKey, false); err != nil {
			return nil, err
		}
		if isNil, err = isNilOrEmpty(result); err != nil {
			return nil, err
		}
		if isNil {
			if cacheGetIfAbsent != nil {
				if err = cacheGetIfAbsent(key); err != nil {
					return nil, err
				}
			}
		}
	}
	return result, nil
}

func (r RedisDistributedCache) defaultLockKey(key string) string {
	appName := base.GetConfig().App.Name
	if appName == "" {
		appName = "default"
	}
	return "lock:" + appName + ":" + key
}

// CheckBloomFilter 检查布隆过滤器中是否存在某个键
// @param ctx 上下文
// @param bloomFilter 布隆过滤器的键
// @param key 待检查的键
// @param exceptKey 失效缓存的键
// @return int 0: 不存在 1: 存在 -1: 布隆过滤器不存在
func (r RedisDistributedCache) CheckBloomFilter(ctx context.Context, bloomFilter string, key string, exceptKey string) (int, error) {
	luaScript := `
  -- 检查布隆过滤器是否存在
  if redis.call("EXISTS", KEYS[1]) == 0 then
    return -1
  end

  -- 检查 exceptKey 是否存在
  if ARGV[1] ~= "" and redis.call("GET", ARGV[1]) then
    return 0
  end

  -- 检查 bloomKey 是否在布隆过滤器中
  if redis.call("BF.EXISTS", KEYS[1], ARGV[2]) == 1 then
    return 1
  end

  return 0
`

	result, err := r.rdb.Eval(ctx, luaScript, []string{bloomFilter}, exceptKey, key).Result()
	if err != nil {
		return 0, err
	}
	return int(result.(int64)), nil
}

func (r RedisDistributedCache) SafePut(
	ctx context.Context,
	key string,
	value interface{},
	expiration time.Duration,
	bloomFilter string,
	bloomKey string,
) error {
	//if err := r.Put(ctx, key, value, expiration); err != nil {
	//	return err
	//}
	//return r.rdb.BFAdd(ctx, bloomFilter, bloomKey).Err()
	luaScript := `
        redis.call("SET", KEYS[1], ARGV[1], "EX", ARGV[2])
        return redis.call("BF.ADD", KEYS[2], ARGV[3])
    `

	valueBytes, err := sonic.Marshal(value)
	if err != nil {
		return err
	}

	// 尝试使用 EVALSHA 执行缓存的脚本
	result, err := r.rdb.Eval(ctx, luaScript, []string{key, bloomFilter}, valueBytes, int(expiration.Seconds()), bloomKey).Result()
	if err != nil {
		return err
	}

	if result.(int64) == 1 {
		return nil
	}
	return errors.New("failed to add key to Bloom filter")
}

func (r RedisDistributedCache) SafeDelete(ctx context.Context, key string, exceptBloomKey string) error {
	if ok, err := r.Delete(ctx, key); err != nil {
		return err
	} else if !ok {
		return nil
	}
	if exceptBloomKey != "" {
		if err := r.Put(ctx, exceptBloomKey, "-", NeverExpire); err != nil {
			return err
		}
	}
	return nil
}

func (r RedisDistributedCache) ExistsInBloomFilter(ctx context.Context, bloomFilter, key, exceptKey string) (bool, error) {
	//if exceptKey != "" {
	//	if _, err := r.rdb.Get(ctx, exceptKey).Result(); err != nil {
	//		if errors.Is(err, redis.Nil) {
	//
	//		} else {
	//			return false, err
	//		}
	//	} else {
	//		// 在失效缓存中，意味着从布隆过滤器中删除了
	//		return false, nil
	//	}
	//}
	//return r.rdb.BFExists(ctx, bloomFilter, key).Result()
	luaScript := `
		if ARGV[1] ~= "" and redis.call("GET", ARGV[1]) then
			return 0
		end
		if redis.call("BF.EXISTS", KEYS[1], ARGV[2]) == 1 then
			return 1
		end
		return 0
	`

	result, err := r.rdb.Eval(ctx, luaScript, []string{bloomFilter}, exceptKey, key).Result()
	if err != nil {
		return false, err
	}

	return result.(int64) == 1, nil
}

func (r RedisDistributedCache) CountExistingKeys(ctx context.Context, keys ...string) (int, error) {
	result, err := r.rdb.Exists(ctx, keys...).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

func (r RedisDistributedCache) loadAndSet(
	ctx context.Context,
	key string,
	cacheLoader Loader,
	expiration time.Duration,
	bloomFilter string,
	bloomKey string,
	safeFlag bool,
) (interface{}, error) {
	result, err := cacheLoader()
	if err != nil {
		return nil, err
	}
	isNil := false
	if isNil, err = isNilOrEmpty(result); err != nil {
		return result, err
	}
	if isNil {
		return nil, err
	}
	if safeFlag {
		if err = r.SafePut(ctx, key, result, expiration, bloomFilter, bloomKey); err != nil {
			return nil, err
		}
	} else {
		if err = r.Put(ctx, key, result, expiration); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (r RedisDistributedCache) DoubleDelete(ctx context.Context, key string, delay time.Duration) error {
	return nil
}

func isNilOrEmpty(v interface{}) (bool, error) {
	if v == nil {
		return true, nil
	}

	val := reflect.ValueOf(v)

	switch val.Kind() {
	case reflect.String:
		return val.String() == "", nil
	case reflect.Slice, reflect.Map, reflect.Array:
		return val.Len() == 0, nil
	case reflect.Ptr:
		if val.IsNil() {
			return true, nil
		}
		// 处理指向其他类型的指针
		elem := val.Elem()
		if elem.Kind() == reflect.Struct {
			// 检查结构体的所有字段
			for i := 0; i < elem.NumField(); i++ {
				field := elem.Field(i)
				if !field.IsZero() {
					return false, nil
				}
			}
			return true, nil // 所有字段都是零值
		}
		// 检查指向基本类型的指针
		return elem.IsZero(), nil
	case reflect.Struct:
		// 检查结构体的所有字段
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if !field.IsZero() {
				return false, nil
			}
		}
		return true, nil // 所有字段都是零值
	default:
		return false, errors.New("unsupported type")
	}
}
