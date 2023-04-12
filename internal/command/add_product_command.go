package command

import (
	"github.com/maze1377/manager-vending-machine/internal/machine"
	"github.com/maze1377/manager-vending-machine/internal/models"
)

type AddProductCommand struct {
	product *models.Product
	uid     string
}

func NewAddProductCommand(uid string, product *models.Product) Command {
	return &AddProductCommand{uid: uid, product: product}
}

func (a *AddProductCommand) Execute(vm machine.VendingState) error {
	return vm.AddItem(a.uid, a.product)
}
