package command

import (
	"github.com/maze1377/manager-vending-machine/internal/machine"
)

type ProcessPaymentCommand struct {
	uid           string
	paymentMethod string
	coin          float32
}

func NewProcessPaymentCommand(uid, paymentMethod string, coin float32) Command {
	return &ProcessPaymentCommand{uid: uid, coin: coin, paymentMethod: paymentMethod}
}

func (p *ProcessPaymentCommand) Execute(vm machine.VendingState) error {
	// todo check paymentMethod and some computation
	err := vm.InsertMoney(p.uid, p.coin)
	// todo handle ErrNotEnoughMoney and return back money
	return err
}
