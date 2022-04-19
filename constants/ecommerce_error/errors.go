package ecommerceerror

import "errors"

var (
	ErrCashierNotFound   = errors.New("Cashier Not Found")
	ErrInvalidParameters = errors.New("Invalid Parameters")
)
