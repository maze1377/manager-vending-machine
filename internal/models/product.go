package models

import "fmt"

type Product struct {
	Name  string
	Price int
}

func (p Product) String() string {
	return fmt.Sprintf("Product{Name:%s,Price:%d}", p.Name, p.Price)
}

func NewProduct(name string, price int) *Product {
	return &Product{Name: name, Price: price}
}
