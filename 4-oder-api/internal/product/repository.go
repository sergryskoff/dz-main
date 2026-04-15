// package реализует репозиторий работы с БД
package product

import (
	"order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	DB *db.Db
}

func NewProductRepository(db *db.Db) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

// Здесь реализуем методы объекта ProductRepositor
func (l *ProductRepository) Create(product *Product) (*Product, error) {
	result := l.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (l *ProductRepository) Update(product *Product) (*Product, error) {
	result := l.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (l *ProductRepository) Delete(id uint) error {
	result := l.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (l *ProductRepository) GetByName(name string) (*Product, error) {
	var product Product
	result := l.DB.First(&product, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (l *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := l.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (l *ProductRepository) GetAll() ([]Product, error) {
	var products []Product
	result := l.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
