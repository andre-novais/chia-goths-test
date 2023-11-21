package core

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	openCart Cart
	orders   []Order
	username string
	email    string
}

func makeAccount(username string, email string) Account {
	return Account{openCart: Cart{}, orders: make([]Order, 0), username: username, email: email}
}
