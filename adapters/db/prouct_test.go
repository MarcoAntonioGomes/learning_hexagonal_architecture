package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func setUp() {

	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)

}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := Db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()

}

func createProduct(db *sql.DB) {

	insert := `INSERT INTO products values("abc", "Product test", 10.0, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()

}