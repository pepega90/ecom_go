package order

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/order/repositories"
	"log"

	"gorm.io/gorm"
)

type gormRepo struct {
	db *gorm.DB
}

func NewOrderGormRepo(db *gorm.DB) repositories.OrderRepository {
	return &gormRepo{db}
}
func (d *gormRepo) GetAll() ([]*domain.Order, error) {
	var listOrder []*domain.Order
	err := d.db.Find(&listOrder).Error
	if err != nil {
		log.Fatal("error from repo: ", err)
		return nil, err
	}
	return listOrder, nil
}
func (d *gormRepo) GetById(id uint) (*domain.Order, error) {
	var order domain.Order
	order.ID = id
	err := d.db.Find(&order).Error
	if err != nil {
		log.Fatal("error from repo: ", err)
		return nil, err
	}
	return &order, nil
}

func (d *gormRepo) GetPreloadData(id uint) (*domain.User, error) {
	var user domain.User
	user.ID = id
	err := d.db.Preload("Carts").Preload("Carts.CartItems").Preload("Carts.CartItems.Product").Find(&user).Error
	if err != nil {
		log.Fatal("error from usecase")
		return nil, err
	}
	return &user, nil
}

func (d *gormRepo) Create(order *domain.Order) (*domain.Order, error) {
	var createOrder domain.Order
	createOrder.Total = order.Total
	createOrder.TransaksiID = order.TransaksiID
	createOrder.PaymentType = order.PaymentType
	createOrder.VaNumber = order.VaNumber
	createOrder.PdfUrl = order.PdfUrl
	createOrder.StatusCode = order.StatusCode
	createOrder.UserID = order.UserID
	err := d.db.Create(&createOrder).Error
	if err != nil {
		log.Fatal("error from repo: ", err)
		return nil, err
	}
	return &createOrder, nil
}

func (d *gormRepo) CheckOrder(order *domain.Order) error {
	order.StatusCode = "200"
	err := d.db.Save(&order).Error
	if err != nil {
		log.Fatal("error from repo: ", err)
		return err
	}
	return nil
}

func (d *gormRepo) Delete(id uint) error {
	var order domain.Order
	order.ID = id
	err := d.db.Unscoped().Delete(&order).Error
	if err != nil {
		log.Fatal("error from repo: ", err)
		return err
	}
	return nil
}
