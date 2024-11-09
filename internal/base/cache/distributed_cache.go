package cache

import (
	"context"
	"reflect"
	"time"
)

type Cache interface {
	// Get retrieves a value from the cache
	Get(ctx context.Context, key string, valueType reflect.Type) (interface{}, error)

	HGet(ctx context.Context, key string, gid string) (string, error)

	HGetAll(ctx context.Context, key string) (map[string]string, error)

	// Put stores a value in the cache
	Put(ctx context.Context, key string, value interface{}, expiration time.Duration) error

	// PutIfAbsent stores values if all keys are absent, returns true if successful
	PutIfAbsent(ctx context.Context, key string, value interface{}) (bool, error)

	// Delete removes a value from the cache
	Delete(ctx context.Context, key string) (bool, error)

	// DeleteMultiple removes multiple values from the cache, returns the number of deleted keys
	DeleteMultiple(ctx context.Context, keys []string) (int, error)

	// HasKey checks if a key exists in the cache
	HasKey(ctx context.Context, key string) (bool, error)

	// HIncrBy increments the integer value of a hash field
	HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error)

	// GetInstance retrieves the cache instance
	GetInstance() interface{}
}

// Loader is a function type that loads a value into the cache
type Loader func() (any, error)

// GetFilter is a function type for filtering cache get operations
// 用于解决布隆过滤器无法删除的问题
//type GetFilter func(*redis.Client, string) (bool, error)

// GetIfAbsent is a function type for handling cache get if absent operations
type GetIfAbsent func(string) error

// DistributedCache interface in Go
type DistributedCache interface {
	Cache

	// SafeGet retrieves a value from the cache in a safe manner, preventing cache penetration, breakdown, and avalanche
	SafeGet(
		ctx context.Context,
		key string,
		valueType reflect.Type,
		cacheLoader Loader,
		expiration time.Duration,
	) (interface{}, error)

	SafeGetWithBloomFilter(
		ctx context.Context,
		key string,
		valueType reflect.Type,
		cacheLoader Loader,
		expiration time.Duration,
		bloomFilter string,
		bloomKey string,
	) (interface{}, error)

	SafeGetWithCacheCheckFilter(
		ctx context.Context,
		key string,
		valueType reflect.Type,
		cacheLoader Loader,
		expiration time.Duration,
		bloomFilter string,
		bloomKey string,
		exceptBloomKey string,
	) (interface{}, error)

	SafeGetWithCacheGetIfAbsent(
		ctx context.Context,
		key string,
		valueType reflect.Type,
		cacheLoader Loader,
		expiration time.Duration,
		bloomFilter string,
		bloomKey string,
		exceptBloomKey string,
		cacheGetIfAbsent GetIfAbsent,
	) (interface{}, error)

	// SafePut stores a value in the cache with a custom expiration and bloom filter
	SafePut(
		ctx context.Context,
		key string,
		value any,
		expiration time.Duration,
		bloomFilter string,
		bloomKey string,
	) error

	SafeDelete(ctx context.Context, key string, exceptBloomKey string) error

	ExistsInBloomFilter(ctx context.Context, bloomFilter, key, exceptKey string) (bool, error)

	// CountExistingKeys counts the number of existing keys
	CountExistingKeys(ctx context.Context, keys ...string) (int, error)

	DoubleDelete(ctx context.Context, key string, delay time.Duration) error
}
