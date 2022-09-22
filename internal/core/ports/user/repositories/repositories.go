package repositories

import "ecom_go/internal/core/domain"

type UserRepository interface {
	Create(cart *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	GetCurrentLoginUser(id uint) (*domain.User, error)
}
