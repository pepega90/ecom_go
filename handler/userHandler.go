package handler

import (
	"ecom_go/models/users"
	"ecom_go/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type userHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *userHandler {
	return &userHandler{db}
}

func (u *userHandler) Register(c *fiber.Ctx) error {
	var inputUser users.UserDTO

	if err := c.BodyParser(&inputUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	user := users.User{
		Name:  inputUser.Name,
		Email: inputUser.Email,
		Phone: inputUser.Phone,
	}
	user.HashPassword(inputUser.Password)

	err := u.db.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Sukser register",
		"data":    user,
	})

}

func (u *userHandler) Login(c *fiber.Ctx) error {
	var userInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	var user users.User

	u.db.Where("email = ?", userInput.Email).First(&user)

	err := user.ComparePassword(userInput.Password)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Password salah",
		})
	}

	token, err := utils.GenerateJWT(strconv.Itoa(int(user.ID)))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sukses login",
		"token":   token,
	})
}
func (u *userHandler) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sukses logout",
	})
}
func (u *userHandler) GetCurrentUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	// decode token
	id, err := utils.ParseJwt(cookie)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	var user users.User

	u.db.Preload("Carts").Preload("Carts.CartItems").Preload("Carts.CartItems.Product").Where("id = ?", id).First(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}
