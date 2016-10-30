package products

import (
	"fmt"
	"strings"
)

var productFactories = make(map[string]Product)

func init() {
	Register("p1", NewP1())
	Register("p2", NewP2())
}

type Product interface {
	Name() string
}

func Register(name string, factory Product) {
	if factory == nil {
		panic(fmt.Sprintf("Product factory %s does not exist.", name))
	}
	_, registered := productFactories[name]
	if registered {
		fmt.Printf("Product factory %s already registered. Ignoring.", name)
	}
	productFactories[name] = factory
}

func GetProduct(productCode string) (Product, error) {
	product, ok := productFactories[strings.ToLower(productCode)]
	if !ok {
		return nil, fmt.Errorf("product not found")
	}

	return product, nil
}