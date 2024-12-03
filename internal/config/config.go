package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/quydmfl/zredis/internal/utils"
	"github.com/spf13/viper"
)

type IConfig interface {
	LoadConfig(path string) (*ServerProperties, error)
}

type ServerProperties struct {
	RunID string
	Bind string
	Port int
	Timeout string
}

func LoadConfig(path string) (*ServerProperties, error) {
	var config ServerProperties

	viper.AddConfigPath(path)
	viper.AddConfigPath("$HOME/.zredis/")
	viper.AddConfigPath(".")
	viper.SetConfigName("redis")
	viper.SetConfigType("toml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})

	// defaults config
	viper.SetDefault("runid", utils.RandomString(10))
	viper.SetDefault("bind", "0.0.0.0")
	viper.SetDefault("port", 6379)
	viper.SetDefault("timeout", "60s")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
