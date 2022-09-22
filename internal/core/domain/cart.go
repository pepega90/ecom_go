package domain

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CarItemID uint     `json:"cartitem_id"`
	CartItems CartItem `json:"cart_items" gorm:"foreignKey:CarItemID"`
	UserID    uint     `json:"user_id"`
}
