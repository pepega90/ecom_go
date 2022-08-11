package handler

import (
	"ecom_go/models/products"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type productHandler struct {
	db *gorm.DB
}

func NewProductHandler(db *gorm.DB) *productHandler {
	return &productHandler{db}
}

func (p *productHandler) GetAllProducts(c *fiber.Ctx) error {
	var listProducts []products.Product
	err := p.db.Find(&listProducts).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": listProducts,
	})
}

func (p *productHandler) GetProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Product dengan id %d tidak ditemukan!. error: %s", id, err),
		})
	}
	var prod products.Product
	prod.ID = uint(id)
	err = p.db.Find(&prod).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Product dengan id %d tidak ditemukan!. error: %s", id, err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": prod,
	})
}

func (p *productHandler) CreateProduct(c *fiber.Ctx) error {
	var productInput products.ProductDTO

	if err := c.BodyParser(&productInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	prod := products.Product{
		Title: productInput.Title,
		Desc:  productInput.Desc,
		Price: productInput.Price,
	}

	err := p.db.Create(&prod).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal membuat product, error: %s", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Sukses buat product",
		"data":    prod,
	})
}

func (p *productHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	var prod products.Product
	prod.ID = uint(id)

	err = p.db.Delete(&prod).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal menghapus product, error: %s", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sukses hapus product",
	})

}
