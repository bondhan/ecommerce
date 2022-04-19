package ecommerceerror

import "errors"

var (
	ErrCashierNotFound   = errors.New("Cashier Not Found")
	ErrCategoryNotFound  = errors.New("Category Not Found")
	ErrPaymentNotFound   = errors.New("Payment Not Found")
	ErrProductNotFound   = errors.New("Product Not Found")
	ErrInvalidParameters = errors.New("Invalid Parameters")
)
