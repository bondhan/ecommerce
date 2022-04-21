package usecase

import "github.com/bondhan/ecommerce/modules/payment/model"

type IPaymentUC interface {
	Create(req model.CreatePaymentReq) (model.CreatePaymentResp, error)
	Update(req model.CreatePaymentUpdate) error
	Delete(id uint) error
	List(req model.PaymentPaginated) (model.ListResponse, error)
	Detail(id uint) (model.Payment, error)
}
