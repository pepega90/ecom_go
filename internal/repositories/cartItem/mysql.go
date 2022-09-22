package cartitem

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/cartItem/repositories"
	"log"

	"gorm.io/gorm"
)

type gormRepo struct {
	db *gorm.DB
}

func NewCartItemGormRepo(db *gorm.DB) repositories.CartItemRepository {
	return &gormRepo{db}
}

func (d *gormRepo) CreateCartItem(cartItem *domain.CartItem) (*domain.CartItem, error) {
	var ci domain.CartItem
	ci.Product = cartItem.Product
	ci.ProductID = cartItem.ProductID
	ci.Qty = cartItem.Qty
	err := d.db.Create(&ci).Error
	if err != nil {
		log.Fatal("error from repository: ", err)
		return nil, err
	}
	return &ci, nil
}

func (d *gormRepo) GetProductById(id uint) (*domain.Product, error) {
	var product domain.Product
	product.ID = id
	err := d.db.Find(&product).Error
	if err != nil {
		log.Fatal("error from repository: ", err)
		return nil, err
	}
	return &product, nil
}
