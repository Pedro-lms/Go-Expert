package database

import (
	"github.com/Pedro-lms/goexpert/9-APIs/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	Db *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{Db: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.Db.Create(product).Error
}

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.Db.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindById(product.ID.String())
	if err != nil {
		return err
	}
	return p.Db.Save(product).Error
}
