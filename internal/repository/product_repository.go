package repository

import (
	"database/sql"

	"github.com/kyndverse/catering-management-app/internal/domain/entity"
)

type ProductRepository interface {
	FindAll(tx *sql.Tx) ([]entity.Product, error)
	Create(tx *sql.Tx, product *entity.Product) error
}
