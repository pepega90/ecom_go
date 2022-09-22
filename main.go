package main

import (
	"ecom_go/helpers"
	"ecom_go/middleware"

	cartItemRepo "ecom_go/internal/repositories/cartItem"
	orderRepo "ecom_go/internal/repositories/order"
	productRepo "ecom_go/internal/repositories/product"
	userRepo "ecom_go/internal/repositories/user"

	useCase "ecom_go/internal/core/usecases"
	cartHandler "ecom_go/internal/handlers/cart"
	cartItemHandler "ecom_go/internal/handlers/cartItem"
	orderHandler "ecom_go/internal/handlers/order"
	productHandler "ecom_go/internal/handlers/product"
	userHandler "ecom_go/internal/handlers/user"
	cartRepo "ecom_go/internal/repositories/cart"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	db := helpers.SetupDB()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	// repository
	userRepository := userRepo.NewUserGormRepo(db)
	productRepository := productRepo.NewProductGormRepo(db)
	cartItemRepository := cartItemRepo.NewCartItemGormRepo(db)
	cartRepository := cartRepo.NewCartGormRepo(db)
	orderRepository := orderRepo.NewOrderGormRepo(db)
	// usecases/service
	userUC := useCase.NewUserUseCase(userRepository)
	app.Use(middleware.IsAuthenticated)
	prodUC := useCase.NewProductUseCase(productRepository)
	cartItemUC := useCase.NewCartItemUseCase(cartItemRepository)
	cartUC := useCase.NewCartUseCase(cartRepository)
	orderUC := useCase.NewOrderUseCase(orderRepository)
	// handler
	userHandler.NewUserHandler(userUC, app)
	productHandler.NewProductHandler(prodUC, app)
	cartItemHandler.NewCartItemHandler(cartItemUC, app)
	cartHandler.NewCartHandler(cartUC, app)
	orderHandler.NewOrderHandler(orderUC, app)

	app.Listen(":8080")
}
