package query

import (
	"github.com/bondhan/ecommerce/domain"
	"github.com/bondhan/ecommerce/modules/order/model"
)

type IOrderQ interface {
	Insert(req model.OrderTotal, data []model.ProductSubTotal) (domain.Orders, error)
	Update(req model.CreateOrderUpdate) error
	Delete(id uint) error
	List(req model.OrderPaginated) ([]model.OrderDetailDB, int64, error)
	Detail(id uint) (domain.Orders, error)
}
