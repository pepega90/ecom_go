package user

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/user/repositories"
	"log"

	"gorm.io/gorm"
)

type gormRepo struct {
	db *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) repositories.UserRepository {
	return &gormRepo{db}
}

func (m *gormRepo) Create(user *domain.User) (*domain.User, error) {
	var createdUser domain.User
	createdUser.Email = user.Email
	createdUser.Name = user.Name
	createdUser.Phone = user.Phone
	createdUser.HashPassword(string(user.Password))
	err := m.db.Create(&user).Error
	if err != nil {
		log.Fatal("error from usecase")
		return nil, err
	}
	return user, nil
}
func (m *gormRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	m.db.Where("email = ?", email).First(&user)
	return &user, nil
}
func (m *gormRepo) GetCurrentLoginUser(id uint) (*domain.User, error) {
	var u domain.User
	m.db.Preload("Carts").Preload("Carts.CartItems").Preload("Carts.CartItems.Product").Where("id = ?", id).First(&u)
	return &u, nil
}
