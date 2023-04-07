package main

import (
	"github.com/maze1377/manager-vending-machine/internal/machine"
	"github.com/spf13/cobra"
)

var machineCmd = &cobra.Command{
	Use:   "machine",
	Short: "run vending machine :)",
	Run:   RunMachine,
}

func init() {
	rootCmd.AddCommand(machineCmd)
}

func RunMachine(_ *cobra.Command, _ []string) {
	vendingMachine := machine.GetInstance()
	vendingMachine.Start()
}
