package command

import "github.com/maze1377/manager-vending-machine/internal/machine"

type Command interface {
	Execute(vm *machine.VendingMachine) error
}
