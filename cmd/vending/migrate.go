package main

import (
	"github.com/maze1377/manager-vending-machine/internal/storage"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate db.",
	Run:   RunMigrate,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func RunMigrate(_ *cobra.Command, _ []string) {
	storage.Migrate()
}
