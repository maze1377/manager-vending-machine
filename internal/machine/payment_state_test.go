package machine

import (
	"testing"

	"github.com/maze1377/manager-vending-machine/internal/models"
)

func TestPaymentState_SelectProduct(t *testing.T) {
	// Create a new VendingMachine instance with some products
	products := []*models.Product{
		{Name: "Coke", Price: 50, Quantity: 5},
		{Name: "Pepsi", Price: 60, Quantity: 3},
		{Name: "Sprite", Price: 40, Quantity: 0},
	}
	vm := NewVendingMachine(products)

	// Set the current state to PaymentState with 100 coins
	vm.setCurrentState(NewPaymentState(vm, 50))

	// Attempt to select a product with insufficient funds
	err := vm.SelectProduct("test", "Pepsi")

	// Verify that the error is returned and the product quantity is not updated
	if err != ErrNotEnoughMoney {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if p, _ := vm.findItem("Pepsi"); p.Quantity != 3 {
		t.Errorf("Product quantity should not have changed. Got %d, expected %d", p.Quantity, 3)
	}

	// Attempt to select a product that is out of stock
	err = vm.SelectProduct("test", "Sprite")

	// Verify that the error is returned and the product quantity is not updated
	if err != ErrProductRunningOut {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if p, _ := vm.findItem("Sprite"); p.Quantity != 0 {
		t.Errorf("Product quantity should not have changed. Got %d, expected %d", p.Quantity, 2)
	}

	// Attempt to select a product with sufficient funds
	err = vm.SelectProduct("test", "Coke")

	// Verify that the product is dispensed and the product quantity and machine coins are updated correctly
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if p, _ := vm.findItem("Coke"); p.Quantity != 4 {
		t.Errorf("Product quantity not updated. Got %d, expected %d", p.Quantity, 4)
	}

	if _, ok := vm.currentState.(*dispenseState); !ok {
		t.Errorf("Current state not updated to DispenseState. Got %T, expected *DispenseState", vm.currentState)
	}
}
