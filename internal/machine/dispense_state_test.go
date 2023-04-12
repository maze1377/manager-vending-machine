package machine

import (
	"testing"

	"github.com/maze1377/manager-vending-machine/internal/models"
)

func TestDispenseState_DispenseProduct(t *testing.T) {
	// Create a new VendingMachine instance with some products
	products := []*models.Product{
		{Name: "Coke", Price: 50, Quantity: 5},
		{Name: "Pepsi", Price: 60, Quantity: 3},
		{Name: "Sprite", Price: 40, Quantity: 2},
	}
	vm := &vendingMachine{products: products}

	// Set the current state to DispenseState with the "Coke" product
	product, _ := vm.findItem("Coke")
	vm.setCurrentState(NewDispenseState(vm, product))

	// Attempt to dispense the "Coke" product
	err := vm.DispenseProduct("test", "Coke")
	// Verify that the product is dispensed and the current state is updated to ReadyState
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if product.Quantity != 5 {
		t.Errorf("Product quantity must not change. Got %d, expected %d", product.Quantity, 4)
	}

	if _, ok := vm.currentState.(*readyState); !ok {
		t.Errorf("Current state not updated to ReadyState. Got %T, expected *ReadyState", vm.currentState)
	}
}
