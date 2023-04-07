package machine

import (
	"reflect"
	"testing"
)

func TestNewReadyState(t *testing.T) {
	machine := GetInstance()
	got := NewReadyState(machine)
	want := &ReadyState{machine: machine}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewReadyState() = %v, want %v", got, want)
	}
}

func TestReadyState_dispenseMoney(t *testing.T) {
	machine := GetInstance()
	init := NewReadyState(machine)
	got := init.dispenseMoney()
	want := &ReadyState{machine: machine}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must not change got %v, want %v", got, want)
	}
}

func TestReadyState_dispenseProduct(t *testing.T) {
	machine := GetInstance()
	init := NewReadyState(machine)
	got := init.dispenseProduct()
	want := &ReadyState{machine: machine}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must not change got %v, want %v", got, want)
	}
}

func TestReadyState_insertMoney(t *testing.T) {
	machine := GetInstance()
	init := NewReadyState(machine)
	got := init.insertMoney()
	want := &PaymentState{machine: machine, coins: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must change got %v, want %v", got, want)
	}
}

func TestReadyState_interactWithMenu(t *testing.T) {
	machine := GetInstance()
	init := NewReadyState(machine)
	got := init.interactWithMenu()
	want := &ReadyState{machine: machine}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must not change got %v, want %v", got, want)
	}
}
