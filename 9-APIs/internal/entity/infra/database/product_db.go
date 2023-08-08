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

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.Db.Limit(limit).Offset((page - 1) * limit).
			Order("created_at" + sort).Find(&products).Error
	} else {
		err = p.Db.Order("created_at" + sort).Find(&products).Error
	}
	return products, err
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

func (p *Product) Delete(id string) error {
	_, err := p.FindById(id)
	if err != nil {
		return err
	}
	return p.Db.Delete(product).Error
}
