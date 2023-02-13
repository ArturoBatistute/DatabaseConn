package product

import (
	"WebStore/database"
	"database/sql"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindAllProducts() []Product {

	db := database.Connect()

	result, err := db.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return productRowsToSlice(result)
}

func Insert(product *Product) {

	db := database.Connect()

	insertStatement, err := db.Prepare("INSERT INTO products(name, description, price, quantity) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insertStatement.Exec(product.Name, product.Description, product.Price, product.Quantity)

	defer db.Close()
}

func productRowsToSlice(productRows *sql.Rows) []Product {

	products := []Product{}

	for productRows.Next() {

		product := Product{}

		err := productRows.Scan(&product.ID, &product.Name, &product.Description, &product.Quantity, &product.Price)
		if err != nil {
			panic(err.Error())
		}

		products = append(products, product)
	}

	return products
}
