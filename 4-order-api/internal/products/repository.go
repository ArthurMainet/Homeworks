package products

import (
	"Email-API/packages/db"
	"log"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.DB
}

func NewProductRepository(db *db.DB) *ProductRepository {
	return &ProductRepository{
		Database: db,
	}
}

func (pr *ProductRepository) Create(product *Product) (*Product, error) {
	result := pr.Database.Create(product)
	if result.Error != nil {
		log.Println("Can't create: ", result.Error)
		return nil, result.Error
	}
	return product, nil
}

func (pr *ProductRepository) Update(product *Product) (*Product, error) {
	result := pr.Database.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		log.Println("Can't update: ", result.Error)
		return nil, result.Error
	}
	return product, nil
}

func (pr *ProductRepository) Delete(id int) error {
	result := pr.Database.Delete(&Product{}, id)
	if result.Error != nil {
		log.Println("Can't create: ", result.Error)
		return result.Error
	}
	return nil
}

func (pr *ProductRepository) GetByID(id int) (*Product, error) {
	var product Product
	result := pr.Database.First(&product, " id = ?", id)
	if result.Error != nil {
		log.Println("Can't update: ", result.Error)
		return nil, result.Error
	}
	return &product, nil
}
