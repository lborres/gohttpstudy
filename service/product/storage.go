package product

import (
	"database/sql"

	"github.com/lborres/gohttpstudy/types"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) GetAllProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func scanRowsIntoProduct(row *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.SellPrice,
		&product.Currency,
		&product.Stock,
		&product.CreateDatetime,
		&product.ModDatetime,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}
