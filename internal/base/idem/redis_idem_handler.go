package idem

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"time"
)

const KeyPrefix = "%s:idem:"

type Handler interface {
	IsMessageBeingConsumed(mid string) bool
	HasMessageBeenConsumed(mid string) bool
	MarkMessageAsConsumed(mid string)
	DeleteFlag(mid string)
}

// idemHandler 消息队列幂等处理器
type idemHandler struct {
	rdb       *redis.Client
	keyPrefix string
}

func NewMessageQueueIdempotencyHandler(appName string, rdb *redis.Client) Handler {
	return idemHandler{
		rdb:       rdb,
		keyPrefix: fmt.Sprintf(KeyPrefix, appName),
	}
}

// IsMessageBeingConsumed 消息是否正在被消费
func (h idemHandler) IsMessageBeingConsumed(mid string) bool {
	key := h.keyPrefix + mid
	ok, err := h.rdb.SetNX(context.Background(), key, "0", 2*time.Second).Result()
	if err != nil {
		return true
	}
	if ok {
		return false
	}
	return true
}

// HasMessageBeenConsumed 消息是否已经被消费
func (h idemHandler) HasMessageBeenConsumed(mid string) bool {
	key := h.keyPrefix + mid
	if val := h.rdb.Get(context.Background(), key).String(); val == "1" {
		return true
	}
	return false
}

// MarkMessageAsConsumed 标记消息为已消费
func (h idemHandler) MarkMessageAsConsumed(mid string) {
	key := h.keyPrefix + mid
	err := h.rdb.Set(context.Background(), key, "1", 0).Err()
	if err != nil {
		slog.Error("mark message as consumed failed", "error", err)
		return
	}
}

// DeleteFlag 如果消息处理遇到异常，删除幂等标记
func (h idemHandler) DeleteFlag(mid string) {
	key := h.keyPrefix + mid
	err := h.rdb.Del(context.Background(), key).Err()
	if err != nil {
		slog.Error("delete flag failed", "error", err)
		return
	}
}
