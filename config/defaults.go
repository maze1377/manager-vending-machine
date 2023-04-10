package config

import (
	"github.com/spf13/viper"
)

func setDefaults(config *viper.Viper) {
	config.SetDefault("log-level", "debug")
	config.SetDefault("addr", "0.0.0.0:10000")
	config.SetDefault("db.debug", true)
}
