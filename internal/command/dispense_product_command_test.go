package command

import (
	"testing"

	mockState "github.com/maze1377/manager-vending-machine/internal/machine/mocks"
)

func TestDispenseProductCommand_Execute(t *testing.T) {
	vm := &mockState.VendingState{}
	uid := "test-id"
	productName := "test"
	vm.On("DispenseProduct", uid, productName).Return(nil).Once()
	cmd := NewDispenseProductCommand(uid, productName)
	err := cmd.Execute(vm)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	vm.AssertExpectations(t)
}
