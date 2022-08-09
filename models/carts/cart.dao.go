package carts

import (
	"ecom_go/models/products"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CarItemID uint     `json:"cartitem_id"`
	CartItems CartItem `json:"cart_items" gorm:"foreignKey:CarItemID"`
	UserID    uint     `json:"user_id"`
}

type CartItem struct {
	gorm.Model
	Qty       int              `json:"qty"`
	ProductID uint             `json:"product_id"`
	Product   products.Product `json:"product" gorm:"foreignKey:ProductID"`
}
