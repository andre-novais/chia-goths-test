package core

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	products []Product
}

func (c Cart) AddProduct(p Product) {
	c.products = append(c.products, p)
}

func (c Cart) IsClosable() bool {
	return len(c.products) > 0
}
