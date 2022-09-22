package repositories

import "ecom_go/internal/core/domain"

type CartItemRepository interface {
	CreateCartItem(cartItem *domain.CartItem) (*domain.CartItem, error)
	GetProductById(id uint) (*domain.Product, error)
}
