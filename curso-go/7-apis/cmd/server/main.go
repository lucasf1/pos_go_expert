package main

import (
	"net/http"

	"github.com/lucasf1/goexpert/7-apis/configs"
	"github.com/lucasf1/goexpert/7-apis/internal/entity"
	"github.com/lucasf1/goexpert/7-apis/internal/infra/database"
	"github.com/lucasf1/goexpert/7-apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", nil)
}
