package config

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BasicConfigChangeListener interface {
	OnConfigChanged()
}

type Configure struct {
	config      *viper.Viper
	LogLevel    string
	Address     string
	MetricAddr  string
	VendingAddr []string
	DebugDb     bool
}

var (
	changeListeners   []BasicConfigChangeListener
	ChangeConfigMutex = &sync.RWMutex{}
	Instance          Configure
)

func NewConfig(configFile string) error {
	config := viper.New()
	setDefaults(config)
	config.SetEnvPrefix("SECRETS")
	config.AutomaticEnv()
	if configFile != "" {
		config.SetConfigFile(configFile)
		err := config.ReadInConfig()
		if err != nil {
			return fmt.Errorf("can't read config file %w", err)
		}
		config.WatchConfig()
		config.OnConfigChange(configChanged)
		AddToChangeListener(&Instance)
	}
	Instance.config = config
	Instance.updateSettings()
	return nil
}

func (br *Configure) updateSettings() {
	br.LogLevel = br.config.GetString("log-level")
	br.Address = br.config.GetString("addr")
	br.DebugDb = br.config.GetBool("db.debug")
	br.MetricAddr = br.config.GetString("metric.addr")
	br.VendingAddr = br.config.GetStringSlice("machine.addr")
}

func (br *Configure) OnConfigChanged() {
	br.updateSettings()
}

func AddToChangeListener(listener BasicConfigChangeListener) {
	ChangeConfigMutex.Lock()
	changeListeners = append(changeListeners, listener)
	ChangeConfigMutex.Unlock()
}

func configChanged(fsnotify.Event) {
	log.Warning("Config Changed.... Reloading Config")
	ChangeConfigMutex.Lock()
	defer ChangeConfigMutex.Unlock()
	for _, listener := range changeListeners {
		listener.OnConfigChanged()
	}
}
