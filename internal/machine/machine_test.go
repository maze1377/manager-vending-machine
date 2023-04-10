package machine

import (
	"testing"

	"github.com/maze1377/manager-vending-machine/internal/models"
)

func TestMachine_session_user(t *testing.T) {
	// Create a new VendingMachine instance with some products
	products := []*models.Product{
		{Name: "Coke", Price: 50, Quantity: 5},
		{Name: "Pepsi", Price: 60, Quantity: 3},
		{Name: "Sprite", Price: 40, Quantity: 2},
	}
	vm := NewVendingMachine(products)

	// Set the current state to ReadyState
	vm.setCurrentState(NewReadyState(vm))

	// Attempt to insert some money into the machine
	err := vm.InsertMoney("test", 100)
	// Verify that the current state is updated to PaymentState and no errors are returned
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	// Now other user insert money and have to get error
	err = vm.InsertMoney("test2", 200)

	if err != ErrMachineBusyNow {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}
