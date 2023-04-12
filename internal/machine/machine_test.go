package machine

import (
	"testing"
)

func TestMachine_session_user(t *testing.T) {
	// Create a new VendingMachine instance with some products
	vm := &vendingMachine{}

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
