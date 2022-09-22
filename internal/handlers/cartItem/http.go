package cartitem

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/cartItem/usecases"

	"github.com/gofiber/fiber/v2"
)

type CartItemHandler struct {
	cartItemUsecase usecases.CartItemUsecase
}

func NewCartItemHandler(cartItemUsecase usecases.CartItemUsecase, r *fiber.App) *CartItemHandler {
	handler := &CartItemHandler{
		cartItemUsecase: cartItemUsecase,
	}
	r.Post("/cart-item", handler.CreateCartItem)
	return handler
}

func (ci *CartItemHandler) CreateCartItem(c *fiber.Ctx) error {
	var cartInput CartItemReq

	if err := c.BodyParser(&cartInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	prod, err := ci.cartItemUsecase.GetProductById(cartInput.ProductID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	item := domain.CartItem{
		Qty:       cartInput.Qty,
		ProductID: cartInput.ProductID,
		Product:   *prod,
	}

	createdCI, err := ci.cartItemUsecase.CreateCartItem(&item)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sukses add to cart",
		"data":    createdCI,
	})
}
