package repositories

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/cart/repositories"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type gormRepo struct {
	db *gorm.DB
}

func NewCartGormRepo(db *gorm.DB) repositories.CartRepository {
	return &gormRepo{db}
}

func (m *gormRepo) Create(cart *domain.Cart) (*domain.Cart, error) {
	var createCart domain.Cart
	createCart.CarItemID = cart.CarItemID
	createCart.CartItems = cart.CartItems
	createCart.UserID = cart.UserID
	err := m.db.Create(&createCart).Error
	if err != nil {
		log.Fatal("error created cart")
		return nil, err
	}
	return &createCart, nil
}
func (m *gormRepo) GetById(id string) (*domain.Cart, error) {
	var cart domain.Cart
	idCart, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	cart.ID = uint(idCart)
	err = m.db.Find(&cart).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &cart, nil

}
func (m *gormRepo) GetByPreload(data string, dest *domain.CartItem) (*domain.CartItem, error) {
	var item *domain.CartItem
	err := m.db.Preload(data).Find(&item).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return item, nil
}
func (m *gormRepo) Delete(id string) (*domain.Cart, error) {
	var delCart domain.Cart
	idCart, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	delCart.ID = uint(idCart)
	err = m.db.Unscoped().Delete(&delCart).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &delCart, nil
}
