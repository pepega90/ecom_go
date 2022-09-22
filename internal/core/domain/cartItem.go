package domain

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Qty       int     `json:"qty"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}
