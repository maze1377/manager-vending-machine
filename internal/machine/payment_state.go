package machine

import "github.com/maze1377/manager-vending-machine/internal/models"

type paymentState struct {
	DefaultBehaviour
	coins int
}

func NewPaymentState(machine *VendingMachine, coins int) State {
	return &paymentState{DefaultBehaviour{machine: machine}, coins}
}

func (p *paymentState) SelectProduct(productName string) error {
	product, err := p.machine.findItem(productName)
	if err != nil {
		return err
	}

	if product.Quantity <= 0 {
		return ErrProductRunningOut
	}
	if product.Price > p.coins { // todo add other type of notification if needed
		return ErrNotEnoughMoney
	}
	product.Quantity--
	p.machine.NotifyObservers(models.Payment, "successful", p.coins)
	p.machine.setCurrentState(NewDispenseState(p.machine, product))
	return nil
}
