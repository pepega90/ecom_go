package repositories

import "ecom_go/internal/core/domain"

type ProductRepository interface {
	GetAll() ([]*domain.Product, error)
	GetProduct(id uint) (*domain.Product, error)
	CreateProduct(prod *domain.Product) (*domain.Product, error)
	DeleteProduct(id uint) error
}
