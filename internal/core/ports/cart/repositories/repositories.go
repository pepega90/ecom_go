package repositories

import "ecom_go/internal/core/domain"

type CartRepository interface {
	Create(cart *domain.Cart) (*domain.Cart, error)
	GetById(id string) (*domain.Cart, error)
	GetByPreload(data string, dest *domain.CartItem) (*domain.CartItem, error)
	Delete(id string) (*domain.Cart, error)
}
