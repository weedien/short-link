package config

import (
	"fmt"
	"github.com/bytedance/sonic"
	"shortlink/internal/base"
)

var config = new(Config)

type Config struct {
	base.Config `mapstructure:",squash"`

	// AppLink 短链接服务配置
	AppLink struct {
		Domain            string   `mapstructure:"domain"`
		Whitelist         []string `mapstructure:"whitelist"`
		DefaultFaviconUrl string   `mapstructure:"default_favicon_url"`
		MaxAttempts       int      `mapstructure:"max_attempts"`
		BaseRoutePrefix   string   `mapstructure:"base_route_prefix"`
		MaxLinksPerGroup  int      `mapstructure:"max_links_per_group"`
		Default           struct {
			Gid        string `mapstructure:"gid"`
			Expiration int    `mapstructure:"expiration"`
		} `mapstructure:"default"`
	} `mapstructure:"app_link"`
}

func Get() Config {
	return *config
}

func Init() {
	base.InitConfig(config)

	configJson, _ := sonic.MarshalIndent(config, "", "  ")

	fmt.Println("config: \n", string(configJson))
}
