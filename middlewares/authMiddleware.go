package middlewares

import (
	"ecom_go/utils"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, err := utils.ParseJwt(cookie)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Maaf, anda harus login terlebih dahulu",
		})
	}
	c.Set("user_id", id)
	return c.Next()
}
