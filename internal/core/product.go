package core

import (
	"strings"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Price int
	Name  string
	Slug  string
	Sku   string
}

func MakeProduct(price int, name string, sku string) Product {
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "_"))

	return Product{Price: price, Name: name, Sku: sku, Slug: slug}
}
