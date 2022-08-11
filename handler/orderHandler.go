package handler

import (
	"ecom_go/models/orders"
	"ecom_go/models/users"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/example"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type orderHandler struct {
	db *gorm.DB
}

func NewOrderHandler(db *gorm.DB) *orderHandler {
	return &orderHandler{db}
}

func (o *orderHandler) CreateSNAP(c *fiber.Ctx) error {
	s := snap.Client{}
	s.New(midtrans.ServerKey, midtrans.Sandbox)

	var user users.User

	// find user
	currentUserID, err := strconv.Atoi(c.GetRespHeaders()["User_id"])
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("User dengan id %d tidak ditemukan. error : %s", currentUserID, err),
		})
	}
	user.ID = uint(currentUserID)
	o.db.Preload("Carts").Preload("Carts.CartItems").Preload("Carts.CartItems.Product").Find(&user)

	var items []midtrans.ItemDetails
	total := 0
	for _, v := range user.Carts {
		prod := midtrans.ItemDetails{
			ID:    strconv.Itoa(int(v.CarItemID)),
			Name:  v.CartItems.Product.Title,
			Price: int64(v.CartItems.Product.Price),
			Qty:   int32(v.CartItems.Qty),
		}
		items = append(items, prod)
		total += v.CartItems.Qty * v.CartItems.Product.Price
	}

	// midtrans payment
	req := &snap.Request{TransactionDetails: midtrans.TransactionDetails{
		OrderID:  "MID-GO-ID-" + example.Random(),
		GrossAmt: int64(total),
	},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			LName: user.Name,
			Email: user.Email,
			Phone: strconv.Itoa(user.Phone),
		},
		EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeBankTransfer,
			snap.PaymentTypeIndomaret,
			snap.PaymentTypeAlfamart,
		},
		Items: &items,
	}

	resp, err := s.CreateTransaction(req)

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (o *orderHandler) GetAllOrders(c *fiber.Ctx) error {
	var listOrders []orders.Order
	err := o.db.Find(&listOrders).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(listOrders)
}

func (o *orderHandler) CreateOrder(c *fiber.Ctx) error {
	var orderDto orders.OrderDTO
	if err := c.BodyParser(&orderDto); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	currentUserID, err := strconv.Atoi(c.GetRespHeaders()["User_id"])
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("User dengan id %d tidak ditemukan. error : %s", currentUserID, err),
		})
	}

	order := orders.Order{
		Total:       int(orderDto.GrossAmt),
		PaymentType: orderDto.PaymentType,
		VaNumber:    orderDto.VaNumber,
		PdfUrl:      orderDto.PdfUrl,
		StatusCode:  orderDto.StatusCode,
		UserID:      uint(currentUserID),
	}
	o.db.Create(&order)

	return c.Status(fiber.StatusOK).JSON(order)
}

func (o *orderHandler) DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	var order orders.Order
	order.ID = uint(id)
	err = o.db.Unscoped().Delete(&order).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menghapus order",
	})
}
