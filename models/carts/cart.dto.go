package carts

type CartDTO struct {
	CartItemID uint `json:"cartitem_id"`
}

type CartItemDTO struct {
	ID        uint `json:"id"`
	Qty       int  `json:"qty"`
	ProductID uint `json:"product_id"`
}
