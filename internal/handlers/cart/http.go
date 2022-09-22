package cart

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/cart/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	cartUseCase usecases.CartUseCase
}

func NewCartHandler(cartUseCase usecases.CartUseCase, r *fiber.App) *CartHandler {
	handler := &CartHandler{
		cartUseCase: cartUseCase,
	}
	r.Post("/carts", handler.Create)
	r.Delete("/carts/:id", handler.Delete)
	return handler
}

func (ct *CartHandler) Create(c *fiber.Ctx) error {
	var cartInput CartReq

	if err := c.BodyParser(&cartInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	var item domain.CartItem
	item.ID = cartInput.CartItemID

	t, err := ct.cartUseCase.GetByPreload("Product", &item)
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

	cart := domain.Cart{
		UserID:    uint(currentLoginUserID),
		CarItemID: cartInput.CartItemID,
		CartItems: *t,
	}

	createdCart, err := ct.cartUseCase.Create(&cart)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sukses",
		"data":    createdCart,
	})
}

func (ct *CartHandler) Delete(c *fiber.Ctx) error {
	_, err := ct.cartUseCase.Delete(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menghapus ranjang!",
	})
}
