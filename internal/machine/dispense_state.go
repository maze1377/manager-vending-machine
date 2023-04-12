package machine

import (
	"fmt"

	"github.com/maze1377/manager-vending-machine/internal/models"
)

type dispenseState struct {
	DefaultBehaviour
	product *models.Product
}

func NewDispenseState(machine *vendingMachine, product *models.Product) State {
	return &dispenseState{DefaultBehaviour{machine: machine}, product}
}

func (d *dispenseState) DispenseProduct(productName string) error {
	if d.product.Name != productName {
		return ErrNotExistForDispense
	}
	d.machine.NotifyObservers(models.Payment, true, fmt.Sprintf("dispense Product:%s", productName))
	d.machine.setCurrentState(NewReadyState(d.machine))
	return nil
}
