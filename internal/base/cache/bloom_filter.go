package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const (
	ShortUriCreateBloomFilter = "shortUriCreateBloomFilter"
	ErrorRate                 = 0.0001
	Capacity                  = 1000_000
)

func setUpBloomFilter(rdb *redis.Client) {
	// 如果布隆过滤器存在则跳过
	exists, err := rdb.Exists(context.Background(), ShortUriCreateBloomFilter).Result()
	if err != nil {
		panic(fmt.Errorf("failed to check bloom filter existence: %v", err))
	}
	if exists == 1 {
		return
	}
	_, err = rdb.BFReserve(context.Background(), ShortUriCreateBloomFilter, ErrorRate, Capacity).Result()
	if err != nil {
		panic(fmt.Errorf("failed to setup bloom filter: %v", err))
	}
}
