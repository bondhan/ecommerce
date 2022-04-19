package presenter

import (
	. "github.com/bondhan/ecommerce/infrastructure"
	"github.com/bondhan/ecommerce/modules/product/model"
	"github.com/bondhan/ecommerce/modules/product/usecase"
	"net/http"
)

type productP struct {
	ProductUC usecase.IProductUC
}

func NewProductP(productUC usecase.IProductUC) IProductP {
	return &productP{
		ProductUC: productUC,
	}
}

func (c *productP) Create(w http.ResponseWriter, r *http.Request) {
	req, err := model.NewProduct(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	resp, err := c.ProductUC.Create(req)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, resp)
}

func (c *productP) Update(w http.ResponseWriter, r *http.Request) {
	req, err := model.UpdateProduct(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	err = c.ProductUC.Update(req)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	Success(w, http.StatusOK)
}

func (c *productP) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetProductID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	err = c.ProductUC.Delete(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	Success(w, http.StatusOK)
}
func (c *productP) List(w http.ResponseWriter, r *http.Request) {
	page, err := model.NewProductPaginatedReq(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}
	res, err := c.ProductUC.List(page)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}
	SuccessJSON(w, http.StatusOK, res)

}

func (c *productP) Detail(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetProductID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.ProductUC.Detail(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)
}
