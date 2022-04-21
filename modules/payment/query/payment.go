package query

import (
	"github.com/bondhan/ecommerce/domain"
	"github.com/bondhan/ecommerce/modules/payment/model"
)

type IPaymentQ interface {
	Insert(req model.CreatePaymentReq) (domain.Payments, error)
	Update(req model.CreatePaymentUpdate) error
	Delete(id uint) error
	List(req model.PaymentPaginated) ([]domain.Payments, int64, error)
	Detail(id uint) (domain.Payments, error)
}
