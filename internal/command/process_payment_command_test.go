package command

import (
	"testing"

	mockState "github.com/maze1377/manager-vending-machine/internal/machine/mocks"
)

func TestProcessPaymentCommand_Execute(t *testing.T) {
	vm := &mockState.VendingState{}
	uid := "test-id"
	paymentMethod := "test"
	coin := float32(3.2)
	vm.On("InsertMoney", uid, coin).Return(nil).Once()
	cmd := NewProcessPaymentCommand(uid, paymentMethod, coin)
	err := cmd.Execute(vm)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	vm.AssertExpectations(t)
}
