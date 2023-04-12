package command

import (
	"github.com/maze1377/manager-vending-machine/internal/machine"
)

type SelectProductCommand struct {
	uid         string
	productName string
}

func NewSelectProductCommand(uid, productName string) Command {
	return &SelectProductCommand{uid: uid, productName: productName}
}

func (s *SelectProductCommand) Execute(vm machine.VendingState) error {
	return vm.SelectProduct(s.uid, s.productName)
}
