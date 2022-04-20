package usecase

import "github.com/bondhan/ecommerce/modules/cashier/model"

type ICashierUC interface {
	Create(req model.CreateCashierReq) (model.CreateCashierResp, error)
	Update(req model.CreateCashierUpdate) error
	Delete(id uint) error
	List(req model.CashierPaginated) (model.ListResponse, error)
	Detail(id uint) (model.Cashier, error)
	PassCode(id uint) (model.Passcode, error)
	Login(id uint, passcode model.CreatePasscodeReq) (model.Token, error)
	Logout(id uint, passcode model.CreatePasscodeReq) error
}
