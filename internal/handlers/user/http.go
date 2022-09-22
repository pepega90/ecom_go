package user

import (
	"ecom_go/helpers"
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/user/usecases"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase, r *fiber.App) *UserHandler {
	handler := &UserHandler{
		userUseCase: userUseCase,
	}
	r.Post("/register", handler.Create)
	r.Post("/login", handler.Login)
	// r.Use(middlewares.IsAuthenticated)
	r.Post("/logout", handler.Logout)
	r.Get("/user", handler.GetCurrentUser)
	return handler
}

func (u *UserHandler) Create(c *fiber.Ctx) error {
	var inputUser UserReq

	if err := c.BodyParser(&inputUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	user := domain.User{
		Name:  inputUser.Name,
		Email: inputUser.Email,
		Phone: inputUser.Phone,
	}
	user.HashPassword(inputUser.Password)
	createdUser, err := u.userUseCase.Create(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Sukser register",
		"data":    createdUser,
	})
}
func (u *UserHandler) Login(c *fiber.Ctx) error {
	var userInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	user, err := u.userUseCase.FindByEmail(userInput.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	err = user.ComparePassword(userInput.Password)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Password salah",
		})
	}

	token, err := helpers.GenerateJWT(strconv.Itoa(int(user.ID)))
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

func (u *UserHandler) GetCurrentUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	// decode token
	id, err := helpers.ParseJwt(cookie)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}
	idUser, _ := strconv.Atoi(id)

	user, _ := u.userUseCase.GetCurrentLoginUser(uint(idUser))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}
func (u *UserHandler) Logout(c *fiber.Ctx) error {
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
