package usecases

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/cart/repositories"
	"ecom_go/internal/core/ports/cart/usecases"
	"log"
)

type cartUseCase struct {
	cartRepo repositories.CartRepository
}

func NewCartUseCase(cartRepo repositories.CartRepository) usecases.CartUseCase {
	return &cartUseCase{
		cartRepo: cartRepo,
	}
}

func (c *cartUseCase) GetByPreload(data string, dest *domain.CartItem) (*domain.CartItem, error) {
	var item *domain.CartItem
	item, err := c.cartRepo.GetByPreload(data, item)
	if err != nil {
		log.Fatal("error from repository", err)
		return nil, err
	}
	return item, nil
}

func (c *cartUseCase) GetById(id string) (*domain.Cart, error) {
	cart, err := c.cartRepo.GetById(id)
	if err != nil {
		log.Fatal("error from repository", err)
		return nil, err
	}
	return cart, nil
}

func (c *cartUseCase) Create(cart *domain.Cart) (*domain.Cart, error) {
	cart, err := c.cartRepo.Create(cart)
	if err != nil {
		log.Fatal("error from repository", err)
		return nil, err
	}
	return cart, nil

}
func (c *cartUseCase) Delete(id string) (*domain.Cart, error) {
	cart, err := c.cartRepo.Delete(id)
	if err != nil {
		log.Fatal("error from repository", err)
		return nil, err
	}
	return cart, nil
}
