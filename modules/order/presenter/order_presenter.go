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
	page, err := model.NewOrderPaginatedReq(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.OrderUC.List(page)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)
}

func (c *orderP) Detail(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetOrderID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.OrderUC.Detail(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)
}

func (c *orderP) Download(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetOrderID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.OrderUC.Download(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)
}

func (c *orderP) DownloadStatus(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetOrderID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.OrderUC.CheckDownload(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)
}

func (c *orderP) Revenues(w http.ResponseWriter, r *http.Request) {
	res, err := c.OrderUC.Revenues()
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)

}

func (c *orderP) Solds(w http.ResponseWriter, r *http.Request) {
	res, err := c.OrderUC.Solds()
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}
	SuccessJSON(w, http.StatusOK, res)
}
