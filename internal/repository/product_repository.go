package repository

import (
	"database/sql"

	"golang-web/internal/domain/entity"
)

type ProductRepository interface {
	FindAll(tx *sql.Tx) ([]entity.Product, error)
	Create(tx *sql.Tx, product *entity.Product) error
}
