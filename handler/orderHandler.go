package handler

import (
	"ecom_go/models/orders"
	"ecom_go/models/users"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type orderHandler struct {
	db *gorm.DB
}

func NewOrderHandler(db *gorm.DB) *orderHandler {
	return &orderHandler{db}
}

func (o *orderHandler) CreateOrder(c *fiber.Ctx) error {

	var order orders.Order
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
	total := 0
	for _, v := range user.Carts {
		total += v.CartItems.Qty * v.CartItems.Product.Price
	}
	order.Total = total
	order.UserID = user.ID
	o.db.Create(&order)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Sukses create order",
		"data":    order,
	})
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
