package machine

import (
	"errors"
	"fmt"

	"github.com/maze1377/manager-vending-machine/internal/models"
)

type readyState struct {
	DefaultBehaviour
}

func NewReadyState(machine *VendingMachine) State {
	return &readyState{
		DefaultBehaviour{machine: machine},
	}
}

func (r *readyState) AddItem(product *models.Product) error {
	foundProduct, err := r.machine.findItem(product.Name)
	if err != nil {
		if !errors.Is(err, ErrProductNotFound) {
			return err
		}
		r.machine.addNewProduct(product)
		return nil
	}
	// If the item already exists, we increase its quantity.
	// Note that we assume the price does not change. If the price changes, we need to rename the item.
	foundProduct.Quantity += product.Quantity
	return nil
}

func (r *readyState) InsertMoney(coin float32) error {
	r.machine.NotifyObservers(models.Payment, true, fmt.Sprintf("number of coin %f", coin))
	r.machine.setCurrentState(NewPaymentState(r.machine, coin))
	return nil
}
