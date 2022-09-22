package usecases

import "ecom_go/internal/core/domain"

type CartItemUsecase interface {
	CreateCartItem(cartItem *domain.CartItem) (*domain.CartItem, error)
	GetProductById(id uint) (*domain.Product, error)
}
