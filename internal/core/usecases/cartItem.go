package usecases

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/cartItem/repositories"
	"ecom_go/internal/core/ports/cartItem/usecases"
	"log"
)

type cartItemUseCase struct {
	cartItemRepository repositories.CartItemRepository
}

func NewCartItemUseCase(cartItemRepository repositories.CartItemRepository) usecases.CartItemUsecase {
	return &cartItemUseCase{
		cartItemRepository: cartItemRepository,
	}
}

func (ci *cartItemUseCase) CreateCartItem(cartItem *domain.CartItem) (*domain.CartItem, error) {
	createdCartItem, err := ci.cartItemRepository.CreateCartItem(cartItem)
	if err != nil {
		log.Fatal("error from repo: ", err)
		return nil, err
	}
	return createdCartItem, nil
}

func (ci *cartItemUseCase) GetProductById(id uint) (*domain.Product, error) {
	product, err := ci.cartItemRepository.GetProductById(id)
	if err != nil {
		log.Fatal("error from repo: ", err)
		return nil, err
	}
	return product, nil
}
