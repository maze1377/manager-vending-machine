package machine

type paymentState struct {
	DefaultBehaviour
	coins float32
}

func NewPaymentState(machine *VendingMachine, coins float32) State {
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
	p.machine.setCurrentState(NewDispenseState(p.machine, product))
	return nil
}
