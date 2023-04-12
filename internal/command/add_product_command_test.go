package command

import (
	"testing"

	mockState "github.com/maze1377/manager-vending-machine/internal/machine/mocks"
	"github.com/maze1377/manager-vending-machine/internal/models"
)

func TestAddProductCommand_Execute(t *testing.T) {
	vm := &mockState.VendingState{}
	uid := "test-id"
	product := models.NewProduct("test", 1.2, 1)
	vm.On("AddItem", uid, product).Return(nil).Once()
	cmd := NewAddProductCommand(uid, product)
	err := cmd.Execute(vm)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	vm.AssertExpectations(t)
}
