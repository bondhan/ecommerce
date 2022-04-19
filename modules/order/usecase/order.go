package usecase

import "github.com/bondhan/ecommerce/modules/order/model"

type IOrderUC interface {
	SubTotal(req []model.SubTotalReq) (model.SubTotal, error)
	Create(req model.OrderReq) (model.OrderTotalResp, error)
	//Update(req model.CreateOrderUpdate) error
	//List(req model.OrderPaginated) (model.ListResponse, error)
	//Detail(id uint) (model.Order, error)
}
