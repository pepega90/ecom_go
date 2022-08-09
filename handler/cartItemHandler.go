package handler

import (
	"ecom_go/models/carts"
	"ecom_go/models/products"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type cartItemHandler struct {
	db *gorm.DB
}

func NewCartItemHandler(db *gorm.DB) *cartItemHandler {
	return &cartItemHandler{db}
}

func (cart *cartItemHandler) CreateCartItem(c *fiber.Ctx) error {
	var cartInput carts.CartItemDTO

	if err := c.BodyParser(&cartInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	var prod products.Product
	prod.ID = cartInput.ProductID

	err := cart.db.Find(&prod).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	item := carts.CartItem{
		Qty:       cartInput.Qty,
		ProductID: cartInput.ProductID,
		Product:   prod,
	}

	err = cart.db.Create(&item).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sukses add to cart",
		"data":    item,
	})
}
