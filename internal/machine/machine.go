package machine

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/maze1377/manager-vending-machine/internal/models"
)

var (
	once           sync.Once
	singleInstance *VendingMachine
)

type VendingMachine struct {
	currentState State
	reader       io.Reader
	writer       io.Writer
	products     []*models.Product
}

func GetInstance() *VendingMachine {
	once.Do(func() {
		singleInstance = &VendingMachine{
			reader: os.Stdin,
			writer: os.Stdout,
			products: []*models.Product{
				models.NewProduct("watter", 1),
				models.NewProduct("soda", 3),
				models.NewProduct("coffee", 2),
			},
		}
		singleInstance.currentState = NewReadyState(singleInstance)
	})
	return singleInstance
}

func (vm *VendingMachine) getProducts() []*models.Product {
	return vm.products
}

func (vm *VendingMachine) println(a ...any) {
	_, _ = fmt.Fprintln(vm.writer, a...)
}

func (vm *VendingMachine) readText() string {
	scanner := bufio.NewScanner(vm.reader)
	scanner.Scan()
	return scanner.Text()
}

func (vm *VendingMachine) Start() {
	vm.println("welcome to VendingMachine")
	for {
		cmds := make(map[string]func() State)
		vm.println("1-insertMoney")
		cmds["1"] = vm.currentState.insertMoney
		vm.println("2-interactWithMenu")
		cmds["2"] = vm.currentState.interactWithMenu
		vm.println("3-dispenseProduct")
		cmds["3"] = vm.currentState.dispenseProduct
		vm.println("4-dispenseMoney")
		cmds["4"] = vm.currentState.dispenseMoney

		input := vm.readText()
		if cmd, ok := cmds[input]; ok {
			vm.currentState = cmd()
		}
	}
}
