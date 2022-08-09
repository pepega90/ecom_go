package main

import (
	"ecom_go/handler"
	"ecom_go/middlewares"
	"ecom_go/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// db
	db := utils.SetupDB()

	// settings cors
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// handler
	userHandler := handler.NewUserHandler(db)
	productHandler := handler.NewProductHandler(db)
	cartHandler := handler.NewCartHandler(db)
	cartItemHandler := handler.NewCartItemHandler(db)
	orderHandler := handler.NewOrderHandler(db)

	// routes
	// users route
	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)
	app.Use(middlewares.IsAuthenticated)
	app.Post("/logout", userHandler.Logout)
	app.Get("/user", userHandler.GetCurrentUser)

	// products
	app.Post("/products", productHandler.CreateProduct)
	app.Delete("/products/:id", productHandler.DeleteProduct)

	// carts
	app.Post("/cart-item", cartItemHandler.CreateCartItem)
	app.Post("/carts", cartHandler.AddToCart)
	app.Delete("/carts/:id", cartHandler.DeleteCart)

	// orders
	app.Post("/create-order", orderHandler.CreateOrder)
	app.Delete("/hapus-order/:id", orderHandler.DeleteOrder)

	// start server
	app.Listen(":8080")
}
