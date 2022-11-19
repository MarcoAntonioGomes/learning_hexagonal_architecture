package db

import (
	"database/sql"
	"learning_hexagonal_architecture/application"

	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	Db *sql.DB
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {

	var product application.Product
	stmt, err := p.Db.Prepare("select id, name, price, status from products where id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
