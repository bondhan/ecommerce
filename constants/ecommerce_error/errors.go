package ecommerceerror

import "errors"

var (
	ErrCashierNotFound   = errors.New("Cashier Not Found")
	ErrCashierNotMatch   = errors.New("Cashier Not Match")
	ErrCategoryNotFound  = errors.New("Category Not Found")
	ErrPaymentNotFound   = errors.New("Payment Not Found")
	ErrProductNotFound   = errors.New("Product Not Found")
	ErrOrderNotFound     = errors.New("Order Not Found")
	ErrInvalidParameters = errors.New("Invalid Parameters")
	ErrOutOfStock        = errors.New("Out of Stock")
	ErrStockChange       = errors.New("Stock has changed")
	ErrPasscodeNotMatch  = errors.New("Passcode Not Match")
	ErrUnauthorized      = errors.New("Unauthorized")
	ErrEmptyBody         = errors.New("EmptyBody")
	ErrEmptyProduct      = errors.New("EmptyProduct")
)
