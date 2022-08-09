package orders

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Total  int  `json:"total"`
	UserID uint `json:"user_id"`
}
