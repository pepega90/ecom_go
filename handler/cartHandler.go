package handler

import (
	"ecom_go/models/carts"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type cartHandler struct {
	db *gorm.DB
}

func NewCartHandler(db *gorm.DB) *cartHandler {
	return &cartHandler{db}
}

func (ct *cartHandler) AddToCart(c *fiber.Ctx) error {
	var cartInput carts.CartDTO

	if err := c.BodyParser(&cartInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	var item carts.CartItem
	item.ID = cartInput.CartItemID

	err := ct.db.Preload("Product").Find(&item).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	currentLoginUserID, err := strconv.Atoi(c.GetRespHeaders()["User_id"])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	cart := carts.Cart{
		UserID:    uint(currentLoginUserID),
		CarItemID: cartInput.CartItemID,
		CartItems: item,
	}

	err = ct.db.Create(&cart).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sukses",
		"data":    cart,
	})
}

func (ct *cartHandler) DeleteCart(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	var ranjang carts.Cart
	ranjang.ID = uint(id)
	err = ct.db.Unscoped().Delete(&ranjang).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menghapus ranjang!",
	})
}
