package usecase

import "github.com/bondhan/ecommerce/modules/order/model"

type IOrderUC interface {
	//Create(req model.CreateOrderReq) (model.CreateOrderResp, error)
	//Update(req model.CreateOrderUpdate) error
	SubTotal(req []model.SubTotalReq) (model.SubTotal, error)
	//List(req model.OrderPaginated) (model.ListResponse, error)
	//Detail(id uint) (model.Order, error)
}
