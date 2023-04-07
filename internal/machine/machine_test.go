package machine

import (
	"reflect"
	"testing"
)

func TestGetInstance_singleton(t *testing.T) {
	got := GetInstance()
	want := GetInstance()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetInstance() = %v, want %v", got, want)
	}
}

func TestVendingMachine_Start(_ *testing.T) {
	// todo input finalize and after that check input
}
