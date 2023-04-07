package config

import (
	"github.com/spf13/viper"
)

func setDefaults(config *viper.Viper) {
	config.SetDefault("log-level", "debug")
}
