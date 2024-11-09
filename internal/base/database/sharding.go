package database

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"gorm.io/sharding"
	"shortlink/internal/base/toolkit"
)

func setupSharding(Db *gorm.DB) {

	shardingByUsername := sharding.Register(sharding.Config{
		ShardingKey:         "username",
		NumberOfShards:      16,
		PrimaryKeyGenerator: sharding.PKSnowflake,
		ShardingAlgorithm:   hashModeShardingAlgorithm(),
		ShardingSuffixs:     shardingSuffix(16),
	}, "user", "group")

	Db.Use(shardingByUsername)

	shardingByGid := sharding.Register(sharding.Config{
		ShardingKey:         "gid",
		NumberOfShards:      16,
		PrimaryKeyGenerator: sharding.PKSnowflake,
		ShardingAlgorithm:   hashModeShardingAlgorithm(),
		ShardingSuffixs:     shardingSuffix(16),
	}, "link")

	Db.Use(shardingByGid)

	shardingByFullShortUrl := sharding.Register(sharding.Config{
		ShardingKey:         "full_short_url",
		NumberOfShards:      16,
		PrimaryKeyGenerator: sharding.PKSnowflake,
		ShardingAlgorithm:   hashModeShardingAlgorithm(),
		ShardingSuffixs:     shardingSuffix(16),
	}, "link_goto")

	Db.Use(shardingByFullShortUrl)
}

func shardingSuffix(numberOfShards int) func() (suffixes []string) {
	return func() (suffixes []string) {
		for i := 0; i < numberOfShards; i++ {
			suffixes = append(suffixes, fmt.Sprintf("_%d", i%numberOfShards))
		}
		return
	}
}

// hash based sharding algorithm
func hashModeShardingAlgorithm() func(columnValue any) (suffix string, err error) {
	return func(value any) (suffix string, err error) {
		if value != "" {
			hashValue := toolkit.SHA256(utils.ToString(value))
			// Convert the first 8 characters of the hash to an integer
			var shard int
			fmt.Sscanf(hashValue[:8], "%x", &shard)
			return fmt.Sprintf("_%d", shard%16), nil
		}
		return "", errors.New("invalid username")
	}
}
