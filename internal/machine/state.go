package machine

type State interface {
	insertMoney() State
	interactWithMenu() State
	dispenseProduct() State
	dispenseMoney() State
}
