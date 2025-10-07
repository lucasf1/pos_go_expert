package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// // Create a new product repository
	// repository := product.NewProductRepository(db)

	// // Create a new product use case
	// useCase := product.NewProductUseCase(repository)

	useCase := NewUseCase(db)

	product, err := useCase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
