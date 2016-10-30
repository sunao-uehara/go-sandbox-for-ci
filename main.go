package main

import (
	"fmt"
	pdt "github.com/sunao-uehara/go-sandbox-for-ci/products"
)

func main() {
	product, err := pdt.GetProduct("p1")
	if err != nil {
		panic("product not found")
	}

	fmt.Printf("%v\n", product)
	product.Name()
}