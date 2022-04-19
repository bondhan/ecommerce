package presenter

import (
	. "github.com/bondhan/ecommerce/infrastructure"
	"github.com/bondhan/ecommerce/modules/order/model"
	"github.com/bondhan/ecommerce/modules/order/usecase"
	"net/http"
)

type orderP struct {
	OrderUC usecase.IOrderUC
}

func NewOrderP(orderUC usecase.IOrderUC) IOrderP {
	return &orderP{
		OrderUC: orderUC,
	}
}

func (c *orderP) Create(w http.ResponseWriter, r *http.Request) {
	req, err := model.NewSOrder(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.OrderUC.Create(req)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)

	SuccessJSON(w, http.StatusOK, struct{}{})
}

func (c *orderP) Update(w http.ResponseWriter, r *http.Request) {
	//req, err := model.UpdateOrder(r)
	//if err != nil {
	//	Error(w, http.StatusBadRequest, err)
	//	return
	//}
	//
	//err = c.OrderUC.Update(req)
	//if err != nil {
	//	Error(w, http.StatusBadRequest, err)
	//	return
	//}

	Success(w, http.StatusOK)
}

func (c *orderP) SubTotal(w http.ResponseWriter, r *http.Request) {
	req, err := model.NewSubtotalOrder(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.OrderUC.SubTotal(req)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)
}

func (c *orderP) List(w http.ResponseWriter, r *http.Request) {

	SuccessJSON(w, http.StatusOK, struct{}{})

}

func (c *orderP) Detail(w http.ResponseWriter, r *http.Request) {

	SuccessJSON(w, http.StatusOK, struct{}{})
}

func (c *orderP) Download(w http.ResponseWriter, r *http.Request) {

	Success(w, http.StatusOK)
}

func (c *orderP) DownloadStatus(w http.ResponseWriter, r *http.Request) {

	Success(w, http.StatusOK)
}
