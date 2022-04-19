package usecase

import "github.com/bondhan/ecommerce/modules/cashier/model"

type ICashierUC interface {
	Create(req model.CreateCashierReq) (model.CreateCashierResp, error)
	Update(req model.CreateCashierUpdate) error
	Delete(id uint) error
	List(req model.CashierPaginated) (model.ListResponse, error)
	Detail(id uint) (model.Cashier, error)
}
