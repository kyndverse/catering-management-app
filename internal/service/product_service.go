package service

import (
	"context"
	"database/sql"

	"github.com/kyndverse/catering-management-app/internal/domain/entity"
	"github.com/kyndverse/catering-management-app/internal/model"
	"github.com/kyndverse/catering-management-app/internal/repository"
	"github.com/kyndverse/catering-management-app/internal/validator"
	"github.com/sirupsen/logrus"
)

type ProductService struct {
	Log               *logrus.Logger
	DB                *sql.DB
	Validator         *validator.Validator
	ProductRepository repository.ProductRepository
}

func NewProductService(logger *logrus.Logger, db *sql.DB, validator *validator.Validator, productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		Log:               logger,
		DB:                db,
		Validator:         validator,
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

func (p *ProductService) AddProduct(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, validator.ValidatorErrors, error) {
	errors := p.Validator.ValidateStruct(request)
	if errors != nil {
		return nil, errors, nil
	}

	tx, err := p.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	product := &entity.Product{
		Nama:       request.Nama,
		Harga:      request.Harga,
		Stok:       request.Stok,
		Keterangan: request.Keterangan,
	}

	if err := p.ProductRepository.Create(tx, product); err != nil {
		return nil, nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return &model.ProductResponse{
		ID:         product.ID,
		Nama:       product.Nama,
		Harga:      product.Harga,
		Stok:       product.Stok,
		Keterangan: product.Keterangan,
	}, nil, nil
}
