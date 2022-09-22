package product

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/product/usecases"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productUseCase usecases.ProductUseCase
}

func NewProductHandler(productUseCase usecases.ProductUseCase, r *fiber.App) *ProductHandler {
	handler := &ProductHandler{
		productUseCase: productUseCase,
	}
	r.Get("/products", handler.GetAllProducts)
	r.Get("/products/:id", handler.GetProduct)
	r.Post("/products", handler.CreateProduct)
	r.Delete("/products/:id", handler.DeleteProduct)
	return handler
}

func (ph *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	listProducts, err := ph.productUseCase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": listProducts,
	})
}
func (ph *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Product dengan id %d tidak ditemukan!. error: %s", id, err),
		})
	}
	prod, err := ph.productUseCase.GetProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Product dengan id %d tidak ditemukan!. error: %s", id, err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": prod,
	})
}
func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var prodInput ProductReq
	if err := c.BodyParser(&prodInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	prod := domain.Product{
		Title: prodInput.Title,
		Desc:  prodInput.Desc,
		Price: prodInput.Price,
	}

	p, err := ph.productUseCase.CreateProduct(&prod)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal membuat product, error: %s", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Sukses buat product",
		"data":    p,
	})
}
func (ph *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Product dengan id %d tidak ditemukan!. error: %s", id, err),
		})
	}
	err = ph.productUseCase.DeleteProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Product dengan id %d tidak ditemukan!. error: %s", id, err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sukses hapus product",
	})
}
