package command

import (
	"github.com/maze1377/manager-vending-machine/internal/machine"
)

type DispenseProductCommand struct {
	uid         string
	productName string
}

func NewDispenseProductCommand(uid, productName string) Command {
	return &DispenseProductCommand{uid: uid, productName: productName}
}

func (d DispenseProductCommand) Execute(vm *machine.VendingMachine) error {
	return vm.DispenseProduct(d.uid, d.productName)
}
