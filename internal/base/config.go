package base

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

var config = new(Config)

func GetConfig() Config {
	return *config
}

type ConfigInterface interface {
	Get() Config
}

func (c Config) Get() Config {
	return c
}

// Config 通用基础配置
type Config struct {
	// App 应用配置s
	App struct {
		Name    string `mapstructure:"name"`
		Author  string `mapstructure:"author"`
		Version string `mapstructure:"version"`
		Desc    string `mapstructure:"desc"`
	} `mapstructure:"app"`

	// Server web服务配置
	Server struct {
		Port        int    `mapstructure:"port"`
		UseSsl      bool   `mapstructure:"use_ssl"`
		BasePath    string `mapstructure:"base_path"`
		MaxRequests int    `mapstructure:"max_requests"`
	} `mapstructure:"server"`

	// Database 数据库配置
	Database struct {
		Dsn            string `mapstructure:"dsn"`
		EnableSharding bool   `mapstructure:"enable_sharding"`
	} `mapstructure:"database"`

	// Redis 缓存配置
	Redis struct {
		Addr     string `mapstructure:"addr"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Db       int    `mapstructure:"db"`
	} `mapstructure:"redis"`

	// RocketMQ 消息队列配置
	RocketMQ struct {
		NameServer    string   `mapstructure:"name_server"`
		Topics        []string `mapstructure:"topics"`
		NameSpace     string   `mapstructure:"namespace"`
		ConsumerGroup string   `mapstructure:"consumer_group"`
		AccessKey     string   `mapstructure:"access_key"`
		SecretKey     string   `mapstructure:"secret_key"`
	} `mapstructure:"rocketmq"`

	// Email 邮件配置
	Email struct {
		SMTPHost  string `mapstructure:"smtp_host"`
		SMTPPort  int    `mapstructure:"smtp_port"`
		Username  string `mapstructure:"username"`
		AuthToken string `mapstructure:"auth_token"`
	} `mapstructure:"email"`
}

const DefaultConfigName = "config"

func initC(c ConfigInterface) {

	configName := os.Getenv("CONFIG_NAME")
	if configName == "" {
		configName = DefaultConfigName
		env := os.Getenv("ENV")
		if env != "" {
			configName = DefaultConfigName + "_" + env
		}
	}

	viper.SetDefault("server.max_requests", 1000)
	viper.SetDefault("server.port", 8080)

	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(c); err != nil {
			slog.Error("unmarshal config error at runtime", "error", err)
		}
		bc := c.Get()
		config = &bc
	})

	if err := viper.ReadInConfig(); err != nil {
		panic("read config error: " + err.Error())
	}

	if err := viper.Unmarshal(c); err != nil {
		panic("unmarshal config error: " + err.Error())
	}
}

func InitConfig(c ConfigInterface) {
	// 初始化配置
	initC(c)

	bc := c.Get()
	config = &bc
}
