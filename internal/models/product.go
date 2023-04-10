package models

import "fmt"

type Product struct {
	Name     string
	Price    int
	Quantity int
}

func (p Product) String() string {
	return fmt.Sprintf("Product{Name:%s,Price:%d}", p.Name, p.Price)
}

func (p Product) Debug() string {
	return fmt.Sprintf("Product{Name:%s,Price:%d,Quantity:%d}", p.Name, p.Price, p.Quantity)
}

func NewProduct(name string, price, quantity int) *Product {
	return &Product{Name: name, Price: price, Quantity: quantity}
}
