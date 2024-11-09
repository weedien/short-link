package lock

import (
	"context"
	"errors"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisLock struct {
	locker *redislock.Client
}

func NewRedisLock(rdb *redis.Client) RedisLock {
	return RedisLock{locker: redislock.New(rdb)}
}

var locks = map[string]*redislock.Lock{}

func (r RedisLock) Acquire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	// 每1000ms重试一次 重试3次
	backoff := redislock.LimitRetry(redislock.LinearBackoff(1000*time.Millisecond), 3)
	lock, err := r.locker.Obtain(ctx, key, expiration, &redislock.Options{
		RetryStrategy: backoff,
	})
	if errors.Is(err, redislock.ErrNotObtained) {
		// 锁被占用 获取失败
		return false, nil
	} else if err != nil {
		// redis 异常
		return false, err
	}
	locks[key] = lock
	return true, nil
}

func (r RedisLock) Release(ctx context.Context, key string) error {
	if err := locks[key].Release(ctx); err != nil {
		// 锁释放失败和其他错误并不需要区分 在上层统一视为一种外部错误
		return err
	}
	delete(locks, key)
	return nil
}

func (r RedisLock) Refresh(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	if err := locks[key].Refresh(ctx, expiration, nil); err != nil {
		return false, err
	}
	return true, nil
}
