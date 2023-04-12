package machine

import "github.com/maze1377/manager-vending-machine/internal/models"

type VendingState interface {
	AddItem(uid string, product *models.Product) error
	SelectProduct(uid, productName string) error
	DispenseProduct(uid, productName string) error
	InsertMoney(uid string, coin float32) error
	GetProducts() []*models.Product
	AddObserver(id string, fn func(event models.Event, date ...interface{}))
	RemoveObserver(id string)
}
