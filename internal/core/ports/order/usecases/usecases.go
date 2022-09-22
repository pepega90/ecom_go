package usecases

import "ecom_go/internal/core/domain"

type OrderUseCase interface {
	GetAll() ([]*domain.Order, error)
	GetById(id uint) (*domain.Order, error)
	Create(order *domain.Order) (*domain.Order, error)
	Delete(id uint) error
	GetPreloadData(id uint) (*domain.User, error)
	CheckOrder(order *domain.Order) error
}
