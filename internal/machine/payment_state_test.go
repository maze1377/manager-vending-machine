package machine

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/maze1377/manager-vending-machine/internal/models"
)

func TestNewPaymentState(t *testing.T) {
	machine := GetInstance()
	got := NewPaymentState(machine)
	want := &PaymentState{machine: machine, coins: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewPaymentState() = %v, want %v", got, want)
	}
}

func TestPaymentState_dispenseMoney(t *testing.T) {
	machine := GetInstance()
	init := NewPaymentState(machine)
	got := init.dispenseMoney()
	want := &ReadyState{machine: machine}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must change got %v, want %v", got, want)
	}
}

func TestPaymentState_dispenseProduct(t *testing.T) {
	machine := GetInstance()
	init := NewPaymentState(machine)
	got := init.dispenseProduct()
	want := &PaymentState{machine: machine, coins: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must not change got %v, want %v", got, want)
	}
}

func TestPaymentState_insertMoney(t *testing.T) {
	machine := GetInstance()
	init := NewPaymentState(machine)
	got := init.insertMoney()
	want := &PaymentState{machine: machine, coins: 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must change got %v, want %v", got, want)
	}
}

func TestPaymentState_interactWithMenu_coin_check(t *testing.T) {
	product := models.NewProduct("coffee", 2)
	machine := &VendingMachine{
		products: []*models.Product{product},
		reader:   strings.NewReader("1"),
		writer:   os.Stdout,
	}
	init := NewPaymentState(machine)
	got := init.interactWithMenu()
	want := &PaymentState{machine: machine, coins: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("state must not change got %v, want %v", got, want)
	}
	machine.reader = strings.NewReader("1")
	init = init.insertMoney()
	got = init.interactWithMenu()
	want2 := &dispenseState{machine: machine, product: product}
	if !reflect.DeepEqual(got, want2) {
		t.Errorf("state must  change got %v, want %v", got, want2)
	}
}
