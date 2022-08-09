package utils

import (
	"ecom_go/models/carts"
	"ecom_go/models/orders"
	"ecom_go/models/products"
	"ecom_go/models/users"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	dsn := "root:@/ecom_go?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database: %s", err)
	}

	err = db.AutoMigrate(&users.User{},
		&products.Product{},
		&orders.Order{},
		&carts.Cart{},
		&carts.CartItem{},
	)

	if err != nil {
		log.Fatalf("error connecting to database: %s", err)
		panic(err)
	}

	return db
}
