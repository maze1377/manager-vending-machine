package machine

import (
	"github.com/maze1377/manager-vending-machine/internal/models"
)

type dispenseState struct {
	machine *VendingMachine
	product *models.Product
}

func (d *dispenseState) insertMoney() State {
	d.machine.println("Please wait, dispensing product...")
	return d
}

func (d *dispenseState) interactWithMenu() State {
	d.machine.println("Please wait, dispensing product...")
	return d
}

func (d *dispenseState) dispenseProduct() State {
	d.machine.println("thanks have good days:)")
	return NewReadyState(d.machine)
}

func (d *dispenseState) dispenseMoney() State {
	d.machine.println("Please wait, dispensing product...")
	return d
}

func NewDispenseState(machine *VendingMachine, product *models.Product) State {
	machine.println("Dispense")
	machine.println(product)
	machine.println("-----------------------")
	return &dispenseState{machine: machine, product: product}
}
