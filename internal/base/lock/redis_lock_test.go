package lock

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func setupTestRedisClient() *redis.Client {
	// Setup a Redis client for testing
	rdb := redis.NewClient(&redis.Options{
		Addr: "193.112.178.249:6379",
	})
	return rdb
}

func TestRedisLock_Acquire(t *testing.T) {
	rdb := setupTestRedisClient()
	locker := NewRedisLock(rdb)
	ctx := context.Background()

	success, err := locker.Acquire(ctx, "test-key", 5*time.Second)
	assert.NoError(t, err)
	assert.True(t, success)

	// Try to acquire the same lock again
	success, err = locker.Acquire(ctx, "test-key", 5*time.Second)
	assert.NoError(t, err)
	assert.False(t, success)
}

func TestRedisLock_Release(t *testing.T) {
	rdb := setupTestRedisClient()
	locker := NewRedisLock(rdb)
	ctx := context.Background()

	_, err := locker.Acquire(ctx, "test-key", 5*time.Second)
	assert.NoError(t, err)

	err = locker.Release(ctx, "test-key")
	assert.NoError(t, err)

	// Try to release the same lock again
	err = locker.Release(ctx, "test-key")
	assert.Error(t, err)
}

func TestRedisLock_Refresh(t *testing.T) {
	rdb := setupTestRedisClient()
	locker := NewRedisLock(rdb)
	ctx := context.Background()

	_, err := locker.Acquire(ctx, "test-key", 5*time.Second)
	assert.NoError(t, err)

	success, err := locker.Refresh(ctx, "test-key", 5*time.Second)
	assert.NoError(t, err)
	assert.True(t, success)
}
