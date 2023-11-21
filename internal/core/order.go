package core

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	checkout Checkout
}
