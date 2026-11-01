package service

import (
	"context"
	"database/sql"

	"github.com/kyndverse/catering-management-app/internal/domain/entity"
	"github.com/kyndverse/catering-management-app/internal/repository"
)

type ProductService struct {
	DB                *sql.DB
	ProductRepository repository.ProductRepository
	// validate
	// logger
}

func NewProductService(db *sql.DB, productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		DB:                db,
		ProductRepository: productRepository,
	}
}

func (p *ProductService) GetAllProduct(ctx context.Context) ([]entity.Product, error) {
	tx, err := p.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	products, err := p.ProductRepository.FindAll(tx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return products, nil
}
