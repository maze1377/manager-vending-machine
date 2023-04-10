package machine

import (
	"testing"

	"github.com/maze1377/manager-vending-machine/internal/models"
)

func TestDefaultBehaviour_state_call(t *testing.T) {
	// Create a new VendingMachine without product
	vm := NewVendingMachine(nil)

	// Set the current state to DefaultBehaviour
	vm.setCurrentState(&DefaultBehaviour{machine: vm})

	// Attempt to add an item to the machine
	err := vm.AddItem("123", &models.Product{Name: "Fanta", Price: 45, Quantity: 1})

	// Verify that the AddItem function returns an error
	if err != ErrTransactionNotValid {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	// Attempt to select an item to the machine
	err = vm.SelectProduct("123", "Fanta")

	// Verify that the select an item function returns an error
	if err != ErrTransactionNotValid {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	// Attempt to dispense an item to the machine
	err = vm.DispenseProduct("123", "Fanta")

	// Verify that the dispense an item function returns an error
	if err != ErrTransactionNotValid {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	// Attempt to insert money to the machine
	err = vm.InsertMoney("123", 84)

	// Verify that the insert money  function returns an error
	if err != ErrTransactionNotValid {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	// Verify that the state not change
	if _, ok := vm.currentState.(*DefaultBehaviour); !ok {
		t.Errorf("Current state should not have changed. Got %T, expected *DefaultBehaviour", vm.currentState)
	}
}
