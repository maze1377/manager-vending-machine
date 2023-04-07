package main

import (
	"github.com/maze1377/manager-vending-machine/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vending <subcommand>",
	Short: "vending machine",
	Run:   nil,
}

var (
	thisInstance = &cmdInstance{}
	configFile   = ""
)

type cmdInstance struct{}

func (c *cmdInstance) adjustLogLevel(level string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.SetLevel(log.DebugLevel)
		log.WithError(err).Error("can't parse log level")
		return
	}
	log.SetLevel(lvl)
	log.Infof("Log level is set to %s", level)
}

func (c *cmdInstance) OnConfigChanged() {
	c.adjustLogLevel(config.Instance.LogLevel)
}

func setupLoggers() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}

func init() {
	setupLoggers()
	rootCmd.PersistentFlags().StringVarP(&configFile, "config-file", "c", configFile, "path to the config file (eg ./config.toml)")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		err := config.NewConfig(configFile)
		if err != nil {
			log.WithError(err).Fatalf("Failed to load the config file (%q)", configFile)
		}
		thisInstance.adjustLogLevel(config.Instance.LogLevel)
		config.AddToChangeListener(thisInstance)
	}
}
