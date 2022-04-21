package presenter

import (
	. "github.com/bondhan/ecommerce/infrastructure"
	"github.com/bondhan/ecommerce/modules/category/model"
	"github.com/bondhan/ecommerce/modules/category/usecase"
	"net/http"
)

type categoryP struct {
	CategoryUC usecase.ICategoryUC
}

func NewCategoryP(categoryUC usecase.ICategoryUC) ICategoryP {
	return &categoryP{
		CategoryUC: categoryUC,
	}
}

func (c *categoryP) Create(w http.ResponseWriter, r *http.Request) {
	req, err := model.NewCategory(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	resp, err := c.CategoryUC.Create(req)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, resp)
}

func (c *categoryP) Update(w http.ResponseWriter, r *http.Request) {
	req, err := model.UpdateCategory(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	err = c.CategoryUC.Update(req)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	Success(w, http.StatusOK)
}

func (c *categoryP) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetCategoryID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	err = c.CategoryUC.Delete(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	Success(w, http.StatusOK)
}
func (c *categoryP) List(w http.ResponseWriter, r *http.Request) {
	page, err := model.NewCategoryPaginatedReq(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}
	res, err := c.CategoryUC.List(page)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}
	SuccessJSON(w, http.StatusOK, res)

}

func (c *categoryP) Detail(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetCategoryID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.CategoryUC.Detail(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)
}
