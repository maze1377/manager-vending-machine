package command

import (
	"testing"

	mockState "github.com/maze1377/manager-vending-machine/internal/machine/mocks"
)

func TestSelectProductCommand_Execute(t *testing.T) {
	vm := &mockState.VendingState{}
	uid := "test-id"
	productName := "test"
	vm.On("SelectProduct", uid, productName).Return(nil).Once()
	cmd := NewSelectProductCommand(uid, productName)
	err := cmd.Execute(vm)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	vm.AssertExpectations(t)
}
