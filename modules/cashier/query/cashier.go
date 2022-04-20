package query

import (
	"github.com/bondhan/ecommerce/domain"
	"github.com/bondhan/ecommerce/modules/cashier/model"
)

type ICashierQ interface {
	Insert(req model.CreateCashierReq) (domain.Cashiers, error)
	Update(req model.CreateCashierUpdate) error
	Delete(id uint) error
	List(req model.CashierPaginated) ([]domain.Cashiers, int64, error)
	Detail(id uint) (domain.Cashiers, error)
	UpdateLogin(id uint, loginStatus string) error
}
