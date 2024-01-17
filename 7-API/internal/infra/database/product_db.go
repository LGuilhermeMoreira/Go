package database

import (
	"API/fundamentos/internal/entity"

	"gorm.io/gorm"
)

type ProductDb struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *ProductDb {
	return &ProductDb{
		DB: db,
	}
}

func (db *ProductDb) Create(product *entity.Product) error {
	return db.DB.Create(product).Error
}

// func (db *ProductDb) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
// 	var products []*entity.Product
// 	var err error

// 	if sort != "" && sort != "desc" && sort != "asc" {
// 		sort = "asc"
// 	}

// 	if limit != 0 && page != 0 {
// 		err = db.DB.Limit(limit).Offset((page - 1) * limit).Order("create_at " + sort).Find(&entity.Product{}).Error
// 	} else {
// 		err = db.DB.Order("create_at " + sort).Find(&entity.Product{}).Error
// 	}

// 	return products, err
// }

func (db *ProductDb) FindAll(sort string) ([]*entity.Product, error) {
	var products []*entity.Product
	var err error

	if sort != "" && sort != "desc" && sort != "asc" {
		sort = "asc"
	}

	err = db.DB.Order("create_at " + sort).Find(&products).Error

	return products, err
}

func (db *ProductDb) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	err := db.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (db *ProductDb) Update(product *entity.Product) error {
	_, err := db.FindById(product.ID.String())

	if err != nil {
		return err
	}
	return db.DB.Save(product).Error
}

func (db *ProductDb) Delete(id string) error {
	_, err := db.FindById(id)

	if err != nil {
		return err
	}
	return db.DB.Delete(&entity.Product{}, "id = ?", id).Error
}
