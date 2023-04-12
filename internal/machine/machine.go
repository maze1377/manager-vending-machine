package machine

import (
	"sync"

	"github.com/maze1377/manager-vending-machine/internal/models"
)

// TODO We need to implement a mechanism to reset the Vending Machine system from a non-idle State to an idle State if no commands are received within a certain time period. This mechanism is commonly referred to as a 'watchdog'."

type vendingMachine struct {
	currentState State
	observers    sync.Map
	uidWorker    string
	products     []*models.Product
	lSession     sync.Mutex
	lState       sync.Mutex
}

func (vm *vendingMachine) canAccessVendingMachine(uid string) bool {
	vm.lSession.Lock()
	defer vm.lSession.Unlock()
	if vm.uidWorker != "" && vm.uidWorker != uid {
		return false
	}
	if vm.uidWorker == "" {
		vm.uidWorker = uid
	}
	return true
}

func (vm *vendingMachine) AddItem(uid string, product *models.Product) error {
	if !vm.canAccessVendingMachine(uid) {
		return ErrMachineBusyNow
	}
	vm.lState.Lock()
	defer vm.lState.Unlock()
	err := vm.currentState.AddItem(product)
	if err == nil {
		vm.lSession.Lock()
		defer vm.lSession.Unlock()
		vm.uidWorker = ""
	}
	return err
}

func (vm *vendingMachine) SelectProduct(uid, productName string) error {
	if !vm.canAccessVendingMachine(uid) {
		return ErrMachineBusyNow
	}
	vm.lState.Lock()
	defer vm.lState.Unlock()
	return vm.currentState.SelectProduct(productName)
}

func (vm *vendingMachine) DispenseProduct(uid, productName string) error {
	if !vm.canAccessVendingMachine(uid) {
		return ErrMachineBusyNow
	}
	vm.lState.Lock()
	defer vm.lState.Unlock()
	err := vm.currentState.DispenseProduct(productName)
	if err == nil {
		vm.lSession.Lock()
		defer vm.lSession.Unlock()
		vm.uidWorker = ""
	}
	return err
}

func (vm *vendingMachine) InsertMoney(uid string, coin float32) error {
	if !vm.canAccessVendingMachine(uid) {
		return ErrMachineBusyNow
	}
	vm.lState.Lock()
	defer vm.lState.Unlock()
	return vm.currentState.InsertMoney(coin)
}

func NewVendingMachine(products []*models.Product) VendingState {
	vm := &vendingMachine{products: products}
	vm.currentState = NewReadyState(vm)
	return vm
}

func (vm *vendingMachine) GetProducts() []*models.Product {
	// maybe we want to isolate VendingMachine so we should copy product list.
	return vm.products
}

func (vm *vendingMachine) findItem(productName string) (*models.Product, error) {
	for _, product := range vm.products {
		if product.Name == productName {
			return product, nil
		}
	}
	return nil, ErrProductNotFound
}

func (vm *vendingMachine) addNewProduct(product *models.Product) {
	vm.products = append(vm.products, product)
}

func (vm *vendingMachine) AddObserver(id string, fn func(event models.Event, date ...interface{})) {
	vm.observers.Store(id, fn)
}

func (vm *vendingMachine) RemoveObserver(id string) {
	vm.observers.Delete(id)
}

func (vm *vendingMachine) NotifyObservers(event models.Event, date ...interface{}) {
	vm.observers.Range(func(key, value interface{}) bool {
		fn := value.(func(event models.Event, date ...interface{}))
		fn(event, date...)
		return true
	})
}

func (vm *vendingMachine) setCurrentState(currentState State) {
	vm.currentState = currentState
}
