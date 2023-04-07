package main

import (
	vending "github.com/maze1377/manager-vending-machine"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version info",
	Run:   logVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func logVersion(_ *cobra.Command, _ []string) {
	log.Infof("vending-version:%s vending-BuildTime:%s vending-Commit: %s", vending.Version, vending.BuildTime, vending.Commit)
}
