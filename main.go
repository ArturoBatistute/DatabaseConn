package main

import (
	product "WebStore/entities"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	product.Insert(&product.Product{Name: "Jacket", Description: "grey", Price: 12, Quantity: 1})
	products := product.FindAllProducts()

	for _, product := range products {

		fmt.Println(product)
	}

	os.Exit(0)
}
