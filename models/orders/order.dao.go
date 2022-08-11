package orders

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Total       int    `json:"total"`
	PaymentType string `json:"payment_type"`
	VaNumber    string `json:"va_number"`
	PdfUrl      string `json:"pdf_url"`
	StatusCode  string `json:"status_code"`
	UserID      uint   `json:"user_id"`
}
