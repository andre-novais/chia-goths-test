package service

import (
	"chia-goths/internal/core"

	"gorm.io/gorm"
)

type ProdutsService struct {
	Db *gorm.DB
}

func (self ProdutsService) List() []core.Product {
	var products []core.Product
	self.Db.Find(&products)
	return products
}

func (self ProdutsService) Create(price int, name string, sku string) core.Product {
	product := core.MakeProduct(price, name, sku)

	self.Db.Create(&product)

	return product
}
