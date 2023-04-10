package machine

import "github.com/maze1377/manager-vending-machine/internal/models"

type State interface {
	AddItem(product *models.Product) error
	SelectProduct(productName string) error
	DispenseProduct(productName string) error
	InsertMoney(coin int) error
}
