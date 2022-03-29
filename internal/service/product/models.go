package product

import "fmt"

type Product struct {
	Name  string
	Price int
}

var allProducts = []Product{
	{Name: "first", Price: 100},
	{Name: "second", Price: 200},
	{Name: "third", Price: 300},
}

func (product *Product) String() string {
	return fmt.Sprintf("Product \"%s\" has price: %d", product.Name, product.Price)
}

func NewProduct(name string, price int) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}
