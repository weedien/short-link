package lock

import (
	"context"
	"time"
)

// DistributedLock is a distributed lock interface
//
// 可以通过 Redis、Zookeeper、Etcd 等实现，虽然采用 Redis 实现的情况较多，
// 但封装一层，而不是直接使用 redislock，可以方便替换其他实现
type DistributedLock interface {
	Acquire(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Release(ctx context.Context, key string) error
	Refresh(ctx context.Context, key string, expiration time.Duration) (bool, error)
}
