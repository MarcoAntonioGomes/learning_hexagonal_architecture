package main

import (
	"database/sql"
	db2 "learning_hexagonal_architecture/adapters/db"
	"learning_hexagonal_architecture/application"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)

	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Save("Product Test", 38)

	productService.Enable(product)
}
