package repository

import (
	"database/sql"

	"github.com/kyndverse/catering-management-app/internal/domain/entity"
	"github.com/sirupsen/logrus"
)

type productRepositoryImpl struct {
	Log *logrus.Logger
}

func NewProductRepository(logger *logrus.Logger) ProductRepository {
	return &productRepositoryImpl{
		Log: logger,
	}
}

func (p *productRepositoryImpl) FindAll(tx *sql.Tx) ([]entity.Product, error) {
	var products []entity.Product

	rows, err := tx.Query("SELECT id, nama, harga, stok, COALESCE(keterangan, '') AS keterangan FROM products WHERE deleted_at IS NULL")
	if err != nil {
		p.Log.WithError(err).Error("error query")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		product := entity.Product{}
		if err := rows.Scan(
			&product.ID,
			&product.Nama,
			&product.Harga,
			&product.Stok,
			&product.Keterangan,
		); err != nil {
			p.Log.WithError(err).Error("error scan")
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		p.Log.WithError(err).Error("error rows")
		return nil, err
	}
	return products, nil
}
