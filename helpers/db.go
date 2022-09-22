package helpers

import (
	"ecom_go/internal/core/domain"
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

	err = db.AutoMigrate(&domain.User{},
		&domain.Product{},
		&domain.Order{},
		&domain.Cart{},
		&domain.CartItem{},
	)

	if err != nil {
		log.Fatalf("error connecting to database: %s", err)
		panic(err)
	}

	return db
}
