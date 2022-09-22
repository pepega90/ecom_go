package product

type ProductReq struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Price int    `json:"price"`
}
