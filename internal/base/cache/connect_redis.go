package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"shortlink/internal/base"
)

var rdb *redis.Client

func ConnectToRedis() *redis.Client {

	config := base.GetConfig().Redis

	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Username: config.Username,
		Password: config.Password,
		DB:       config.Db,
	})

	// 检测连接是否成功
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("failed to connect to redis: %v", err))
	}

	// 初始化布隆过滤器
	setUpBloomFilter(rdb)

	return rdb
}
