package main

import (
	"github.com/maze1377/manager-vending-machine/internal/machine"
	"github.com/maze1377/manager-vending-machine/internal/models"
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
	// todo read product from configmap
	products := []*models.Product{
		{Name: "Coke", Price: 50, Quantity: 5},
		{Name: "Pepsi", Price: 60, Quantity: 3},
		{Name: "Sprite", Price: 40, Quantity: 0},
	}
	_ = machine.NewVendingMachine(products)
	// todo write console panel for vendingMachine
}
