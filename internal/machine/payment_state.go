package machine

import (
	"fmt"
	"strconv"
)

type PaymentState struct {
	machine *VendingMachine
	coins   int
}

func (p *PaymentState) insertMoney() State {
	numberOfCoins := p.coins + 1
	p.machine.println("new coin inserted")
	p.machine.println(fmt.Sprintf("%d coin inserted", numberOfCoins))
	p.machine.println("-----------------------")
	return &PaymentState{p.machine, numberOfCoins}
}

func (p *PaymentState) interactWithMenu() State {
	products := p.machine.getProducts()
	p.machine.println("Menu:")
	for i, product := range products {
		p.machine.println(fmt.Sprintf("%d:%s", i+1, product))
	}

	input := p.machine.readText()
	num, err := strconv.Atoi(input)
	if err != nil {
		p.machine.println("Please input the number of product you want")
		return p
	}
	if num > len(products) || num < 1 {
		p.machine.println("invalid input number")
		return p
	}
	product := products[num-1]
	if product.Price > p.coins {
		p.machine.println("not enough coin")
		return p
	}
	// TODO handle more coin than price product
	return NewDispenseState(p.machine, product)
}

func (p *PaymentState) dispenseProduct() State {
	p.machine.println("Please select an product first")
	return p
}

func (p *PaymentState) dispenseMoney() State {
	p.machine.println("Please dispense your coins")
	return NewReadyState(p.machine)
}

func NewPaymentState(machine *VendingMachine) State {
	machine.println("Payment")
	machine.println("1 coin inserted")
	machine.println("-----------------------")
	return &PaymentState{machine, 1}
}
