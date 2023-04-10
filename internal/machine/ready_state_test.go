package machine

import (
	"testing"

	"github.com/maze1377/manager-vending-machine/internal/models"
)

func TestReadyState_AddItem(t *testing.T) {
	// Create a new VendingMachine instance with some products
	products := []*models.Product{
		{Name: "Coke", Price: 50, Quantity: 5},
		{Name: "Pepsi", Price: 60, Quantity: 3},
		{Name: "Sprite", Price: 40, Quantity: 2},
	}
	vm := NewVendingMachine(products)

	// Set the current state to ReadyState
	vm.setCurrentState(NewReadyState(vm))

	// Attempt to add a new product to the machine
	newProduct := &models.Product{Name: "Fanta", Price: 45, Quantity: 1}
	err := vm.AddItem("test", newProduct)
	// Verify that the new product is added to the machine and the quantity is correct
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if len(vm.products) != 4 {
		t.Errorf("Product count not updated. Got %d, expected %d", len(vm.products), 4)
	}

	if vm.products[3].Name != newProduct.Name || vm.products[3].Price != newProduct.Price || vm.products[3].Quantity != newProduct.Quantity {
		t.Errorf("Product details not updated correctly. Got %v, expected %v", vm.products[3], newProduct)
	}

	// Attempt to add an existing product to the machine
	existingProduct := &models.Product{Name: "Coke", Price: 50, Quantity: 2}
	err = vm.AddItem("test", existingProduct)

	// Verify that the existing product quantity is updated correctly
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if len(vm.products) != 4 {
		t.Errorf("Product count should not have changed. Got %d, expected %d", len(vm.products), 4)
	}

	if vm.products[0].Quantity != 7 {
		t.Errorf("Product quantity not updated correctly. Got %d, expected %d", vm.products[0].Quantity, 7)
	}
}

func TestReadyState_InsertMoney(t *testing.T) {
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
	if _, ok := vm.currentState.(*paymentState); !ok {
		t.Errorf("Current state not updated to PaymentState. Got %T, expected *PaymentState", vm.currentState)
	}
}
