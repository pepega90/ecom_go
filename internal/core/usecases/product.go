package usecases

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/product/repositories"
	"ecom_go/internal/core/ports/product/usecases"
	"log"
)

type productUsecase struct {
	productRepo repositories.ProductRepository
}

func NewProductUseCase(productRepo repositories.ProductRepository) usecases.ProductUseCase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (pu *productUsecase) GetAll() ([]*domain.Product, error) {
	listProduct, err := pu.productRepo.GetAll()
	if err != nil {
		log.Fatal("error from repo", err)
		return nil, err
	}
	return listProduct, nil
}

func (pu *productUsecase) GetProduct(id uint) (*domain.Product, error) {
	prod, err := pu.productRepo.GetProduct(id)
	if err != nil {
		log.Fatal("error from repository", err)
		return nil, err
	}
	return prod, nil
}
func (pu *productUsecase) CreateProduct(prod *domain.Product) (*domain.Product, error) {
	createdProd, err := pu.productRepo.CreateProduct(prod)
	if err != nil {
		log.Fatal("error from repository", err)
		return nil, err
	}
	return createdProd, nil
}
func (pu *productUsecase) DeleteProduct(id uint) error {
	err := pu.productRepo.DeleteProduct(id)
	if err != nil {
		log.Fatal("error from repository", err)
		return err
	}
	return nil
}
