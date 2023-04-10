package machine

import (
	"github.com/maze1377/manager-vending-machine/internal/models"
)

type DefaultBehaviour struct {
	machine *VendingMachine
}

func (d *DefaultBehaviour) AddItem(_ *models.Product) error {
	return ErrTransactionNotValid
}

func (d *DefaultBehaviour) SelectProduct(_ string) error {
	return ErrTransactionNotValid
}

func (d *DefaultBehaviour) DispenseProduct(_ string) error {
	return ErrTransactionNotValid
}

func (d *DefaultBehaviour) InsertMoney(_ float32) error {
	return ErrTransactionNotValid
}
