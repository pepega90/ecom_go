package usecases

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/user/repositories"
	"ecom_go/internal/core/ports/user/usecases"
	"log"
)

type userUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) usecases.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) Create(user *domain.User) (*domain.User, error) {
	user, err := uc.userRepo.Create(user)
	if err != nil {
		log.Fatal("error from repo")
		return nil, err
	}
	return user, nil
}
func (uc *userUseCase) FindByEmail(email string) (*domain.User, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		log.Fatal("error from repo")
		return nil, err
	}
	return user, nil
}
func (uc *userUseCase) GetCurrentLoginUser(id uint) (*domain.User, error) {
	u, err := uc.userRepo.GetCurrentLoginUser(id)
	if err != nil {
		log.Fatal("error from repo")
		return nil, err
	}
	return u, nil
}
