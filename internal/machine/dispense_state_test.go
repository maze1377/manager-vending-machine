package machine

import (
	"reflect"
	"testing"

	"github.com/maze1377/manager-vending-machine/internal/models"
)

func TestNewDispenseState(t *testing.T) {
	machine := GetInstance()
	product := models.NewProduct("test", 10)
	got := NewDispenseState(machine, product)
	want := &dispenseState{machine: machine, product: product}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewDispenseState() = %v, want %v", got, want)
	}
}

func Test_dispenseState_dispenseMoney(t *testing.T) {
	machine := GetInstance()
	product := models.NewProduct("test", 10)
	init := NewDispenseState(machine, product)
	got := init.dispenseMoney()
	want := &dispenseState{machine: machine, product: product}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must not change got %v, want %v", got, want)
	}
}

func Test_dispenseState_dispenseProduct(t *testing.T) {
	machine := GetInstance()
	product := models.NewProduct("test", 10)
	init := NewDispenseState(machine, product)
	got := init.dispenseProduct()
	want := &ReadyState{machine: machine}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must change to ready got %v, want %v", got, want)
	}
}

func Test_dispenseState_insertMoney(t *testing.T) {
	machine := GetInstance()
	product := models.NewProduct("test", 10)
	init := NewDispenseState(machine, product)
	got := init.insertMoney()
	want := &dispenseState{machine: machine, product: product}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must not change got %v, want %v", got, want)
	}
}

func Test_dispenseState_interactWithMenu(t *testing.T) {
	machine := GetInstance()
	product := models.NewProduct("test", 10)
	init := NewDispenseState(machine, product)
	got := init.interactWithMenu()
	want := &dispenseState{machine: machine, product: product}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must not change got %v, want %v", got, want)
	}
}
