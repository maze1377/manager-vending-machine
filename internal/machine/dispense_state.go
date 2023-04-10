package machine

import "github.com/maze1377/manager-vending-machine/internal/models"

type dispenseState struct {
	DefaultBehaviour
	product *models.Product
}

func NewDispenseState(machine *VendingMachine, product *models.Product) State {
	return &dispenseState{DefaultBehaviour{machine: machine}, product}
}

func (d *dispenseState) DispenseProduct(productName string) error {
	if d.product.Name != productName {
		return ErrNotExistForDispense
	}
	d.machine.NotifyObservers(models.Dispensed, d.product)
	d.machine.setCurrentState(NewReadyState(d.machine))
	return nil
}
