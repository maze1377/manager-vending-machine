package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestConfigOperations(t *testing.T) {
	t.Run("update Config", func(t *testing.T) {
		config := viper.New()
		instance := Configure{config: config}
		config.SetDefault("log-level", "debug")
		instance.OnConfigChanged()
		if instance.LogLevel != "debug" {
			t.Error("config not loaded")
		}
	})
}
