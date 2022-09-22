package usecases

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/order/repositories"
	"ecom_go/internal/core/ports/order/usecases"
	"log"
)

type orderUsecase struct {
	orderRepo repositories.OrderRepository
}

func NewOrderUseCase(orderRepo repositories.OrderRepository) usecases.OrderUseCase {
	return &orderUsecase{
		orderRepo: orderRepo,
	}
}

func (oc *orderUsecase) GetAll() ([]*domain.Order, error) {
	listOrder, err := oc.orderRepo.GetAll()
	if err != nil {
		log.Fatal("error from repo: ", err)
		return nil, err
	}
	return listOrder, nil
}
func (oc *orderUsecase) GetById(id uint) (*domain.Order, error) {
	order, err := oc.orderRepo.GetById(id)
	if err != nil {
		log.Fatal("error from repo: ", err)
		return nil, err
	}
	return order, nil
}
func (oc *orderUsecase) Create(order *domain.Order) (*domain.Order, error) {
	created, err := oc.orderRepo.Create(order)
	if err != nil {
		log.Fatal("error from repo: ", err)
		return nil, err
	}
	return created, nil
}

func (oc *orderUsecase) GetPreloadData(id uint) (*domain.User, error) {
	user, err := oc.orderRepo.GetPreloadData(id)
	if err != nil {
		log.Fatal("error from repo: ", err)
		return nil, err
	}
	return user, nil
}

func (oc *orderUsecase) CheckOrder(order *domain.Order) error {
	err := oc.orderRepo.CheckOrder(order)
	if err != nil {
		log.Fatal("error from repo: ", err)
		return err
	}
	return nil
}

func (oc *orderUsecase) Delete(id uint) error {
	err := oc.orderRepo.Delete(id)
	if err != nil {
		log.Fatal("error from repo: ", err)
		return err
	}
	return nil
}
