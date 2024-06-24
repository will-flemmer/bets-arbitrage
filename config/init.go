package config

import "github.com/spf13/viper"

func Init() {
	viper.SetDefault("ENV", "development")
}
