package models

import "fmt"

type Product struct {
	Name     string
	Price    float32
	Quantity int32
}

func (p Product) String() string {
	return fmt.Sprintf("Product{Name:%s,Price:%f}", p.Name, p.Price)
}

func (p Product) Debug() string {
	return fmt.Sprintf("Product{Name:%s,Price:%f,Quantity:%d}", p.Name, p.Price, p.Quantity)
}

func NewProduct(name string, price float32, quantity int32) *Product {
	return &Product{Name: name, Price: price, Quantity: quantity}
}
