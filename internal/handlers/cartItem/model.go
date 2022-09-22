package cartitem

type CartItemReq struct {
	ID        uint `json:"id"`
	Qty       int  `json:"qty"`
	ProductID uint `json:"product_id"`
}
