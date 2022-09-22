package product

import (
	"ecom_go/internal/core/domain"
	"ecom_go/internal/core/ports/product/repositories"
	"log"

	"gorm.io/gorm"
)

type productMysqlRepo struct {
	db *gorm.DB
}

func NewProductGormRepo(db *gorm.DB) repositories.ProductRepository {
	return &productMysqlRepo{
		db: db,
	}
}

func (pm *productMysqlRepo) GetAll() ([]*domain.Product, error) {
	var listProduct []*domain.Product
	err := pm.db.Find(&listProduct).Error
	if err != nil {
		log.Fatal("error : ", err)
		return nil, err
	}
	return listProduct, nil
}
func (pm *productMysqlRepo) GetProduct(id uint) (*domain.Product, error) {
	var product domain.Product
	product.ID = id
	err := pm.db.Find(&product).Error
	if err != nil {
		log.Fatal("error : ", err)
		return nil, err
	}
	return &product, nil
}
func (pm *productMysqlRepo) CreateProduct(prod *domain.Product) (*domain.Product, error) {
	var createdProd domain.Product
	createdProd.Title = prod.Title
	createdProd.Price = prod.Price
	createdProd.Desc = prod.Desc
	err := pm.db.Create(&createdProd).Error
	if err != nil {
		log.Fatal("error : ", err)
		return nil, err
	}
	return &createdProd, nil
}
func (pm *productMysqlRepo) DeleteProduct(id uint) error {
	var prod domain.Product
	prod.ID = id
	err := pm.db.Delete(&prod).Error
	if err != nil {
		log.Fatal("error : ", err)
		return err
	}
	return nil
}
