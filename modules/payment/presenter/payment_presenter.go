package presenter

import (
	. "github.com/bondhan/ecommerce/infrastructure"
	"github.com/bondhan/ecommerce/modules/payment/model"
	"github.com/bondhan/ecommerce/modules/payment/usecase"
	"net/http"
)

type paymentP struct {
	PaymentUC usecase.IPaymentUC
}

func NewPaymentP(paymentUC usecase.IPaymentUC) IPaymentP {
	return &paymentP{
		PaymentUC: paymentUC,
	}
}

func (c *paymentP) Create(w http.ResponseWriter, r *http.Request) {
	req, err := model.NewPayment(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	resp, err := c.PaymentUC.Create(req)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, resp)
}

func (c *paymentP) Update(w http.ResponseWriter, r *http.Request) {
	req, err := model.UpdatePayment(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	err = c.PaymentUC.Update(req)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	Success(w, http.StatusOK)
}

func (c *paymentP) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetPaymentID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	err = c.PaymentUC.Delete(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	Success(w, http.StatusOK)
}
func (c *paymentP) List(w http.ResponseWriter, r *http.Request) {
	page, err := model.NewPaymentPaginatedReq(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}
	res, err := c.PaymentUC.List(page)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}
	SuccessJSON(w, http.StatusOK, res)

}

func (c *paymentP) Detail(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetPaymentID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.PaymentUC.Detail(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)
}
