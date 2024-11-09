package base_config

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type configWithDefault struct {
	key          string
	defaultValue string
}

var (
	RocketMqNameSrv       = configWithDefault{"ROCKETMQ_NAMESRV", "localhost:9876"}
	RocketMqTopic         = configWithDefault{"ROCKETMQ_TOPIC", "shortlink"}
	RocketMqNameSpace     = configWithDefault{"ROCKETMQ_NAMESPACE", ""}
	RocketMqConsumerGroup = configWithDefault{"ROCKETMQ_CONSUMER_GROUP", "shortlink"}
	RocketMqAccessKey     = configWithDefault{"ROCKETMQ_ACCESS_KEY", ""}
	RocketMqSecretKey     = configWithDefault{"ROCKETMQ_SECRET_KEY", ""}
	DSN                   = configWithDefault{"DSN", "host=remote user=weedien password=031209 dbname=wespace search_path=link port=5432 sslmode=disable TimeZone=Asia/Shanghai"}
	RedisAddr             = configWithDefault{"REDIS_ADDR", "localhost:6379"}
	RedisPassword         = configWithDefault{"REDIS_PASSWORD", ""}
	RedisDB               = configWithDefault{"REDIS_DB", "0"}
	EnableSharding        = configWithDefault{"ENABLE_SHARDING", "false"}
	BaseRoutePrefix       = configWithDefault{"BASE_ROUTE_PREFIX", "/api/short-link/v1"}
	Port                  = configWithDefault{"PORT", "8080"}
	LinkDomain            = configWithDefault{"LINK_DOMAIN", "http://localhost:8080"}
	DomainWhiteList       = configWithDefault{"DOMAIN_WHITELIST", ""}
	DefaultFavicon        = configWithDefault{"DEFAULT_FAVICON_URL", "https://cdn.jsdelivr.net/gh/weedien/shortlink@main/static/favicon.ico"}
	UseSSL                = configWithDefault{"USE_SSL", "false"}
	MaxAttempts           = configWithDefault{"MAX_ATTEMPTS", "10"}
)

func (c configWithDefault) String() string {
	return Default(c.key, c.defaultValue)
}

func (c configWithDefault) Int() int {
	value := Default(c.key, c.defaultValue)
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Warn(fmt.Sprintf("Config key %s value %s is not a int", c.key, value))
		return 0
	}
	return i
}

func (c configWithDefault) Bool() bool {
	value := Default(c.key, c.defaultValue)
	b, err := strconv.ParseBool(value)
	if err != nil {
		log.Warn(fmt.Sprintf("Config key %s value %s is not a bool", c.key, value))
		return false
	}
	return b
}

func (c configWithDefault) Array() []string {
	value := Default(c.key, c.defaultValue)
	// split by comma
	return strings.Split(value, ",")
}

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

// Default func to get env value with default value
func Default(key string, defaultValue string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return defaultValue
}

// DefaultInt func to get env value with default value
func DefaultInt(key string, defaultValue int) int {
	c := Config(key)
	if c == "" {
		return defaultValue
	}
	ci, err := strconv.Atoi(c)
	if err != nil {
		return ci
	}
	return defaultValue
}

func DefaultBool(key string, defaultValue bool) bool {
	c := Config(key)
	if c == "" {
		return defaultValue
	}
	cb, err := strconv.ParseBool(c)
	if err != nil {
		return cb
	}
	return defaultValue
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	if err := viper.Unmarshal(config); err != nil {
	//		panic(err)
	//	}
	//})
	//
	//if err := viper.ReadInConfig(); err != nil {
	//	panic(err)
	//}
	//
	//if err := viper.Unmarshal(config); err != nil {
	//	panic(err)
	//}
}
