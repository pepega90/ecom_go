package order

type OrderReq struct {
	GrossAmt    int    `json:"gross_amount"`
	PaymentType string `json:"payment_type"`
	OrderID     string `json:"order_id"`
	VaNumber    string `json:"va_number"`
	PdfUrl      string `json:"pdf_url"`
	StatusCode  string `json:"status_code"`
}
