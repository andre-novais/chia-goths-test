package core

import (
	"gorm.io/gorm"
)

type Checkout struct {
	gorm.Model
	cart          Cart
	paymentMethod int
	address       string
	date          string
	status        int
}

func (c Checkout) IsClosable() bool {
	return c.cart.IsClosable() && c.address != "" && c.date != ""
}

func (c Checkout) MakeOrder() Order {
	c.status = CheckoutClosed
	return Order{checkout: c}
}

const (
	CreditCardPaymentMethod = iota
	DebitCardPaymentMethod
)

const (
	CheckoutOpen = iota
	CheckoutClosed
)
