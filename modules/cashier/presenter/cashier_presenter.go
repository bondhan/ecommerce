package presenter

import (
	"errors"
	ecommerceerror "github.com/bondhan/ecommerce/constants/ecommerce_error"
	. "github.com/bondhan/ecommerce/infrastructure"
	"github.com/bondhan/ecommerce/modules/cashier/model"
	"github.com/bondhan/ecommerce/modules/cashier/usecase"
	"net/http"
)

type cashierP struct {
	CashierUC usecase.ICashierUC
}

func NewCashierP(cashierUC usecase.ICashierUC) ICashierP {
	return &cashierP{
		CashierUC: cashierUC,
	}
}

func (c *cashierP) Create(w http.ResponseWriter, r *http.Request) {
	req, err := model.NewCashier(r)
	if err != nil {
		if errors.Is(err, ecommerceerror.ErrEmptyBody) {
			Error(w, http.StatusBadRequest, err)
			return
		}
		Error(w, http.StatusBadRequest, err)
		return
	}

	resp, err := c.CashierUC.Create(req)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, resp)
}

func (c *cashierP) Update(w http.ResponseWriter, r *http.Request) {
	req, err := model.UpdateCashier(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	err = c.CashierUC.Update(req)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	Success(w, http.StatusOK)
}

func (c *cashierP) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetCashierID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	err = c.CashierUC.Delete(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	Success(w, http.StatusOK)
}
func (c *cashierP) List(w http.ResponseWriter, r *http.Request) {
	page, err := model.NewCashierPaginatedReq(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}
	res, err := c.CashierUC.List(page)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}
	SuccessJSON(w, http.StatusOK, res)

}

func (c *cashierP) Detail(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetCashierID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.CashierUC.Detail(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)
}

func (c *cashierP) PassCode(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetCashierID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.CashierUC.PassCode(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)
}

func (c *cashierP) Login(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetCashierID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	req, err := model.NewPasscode(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.CashierUC.Login(id, req)
	if err != nil {
		if errors.Is(err, ecommerceerror.ErrPasscodeNotMatch) {
			Error(w, http.StatusUnauthorized, err)
			return
		}
		Error(w, http.StatusInternalServerError, err)
		return
	}

	SuccessJSON(w, http.StatusOK, res)
}

func (c *cashierP) Logout(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetCashierID(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	req, err := model.NewPasscode(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	err = c.CashierUC.Logout(id, req)
	if err != nil {
		if errors.Is(err, ecommerceerror.ErrPasscodeNotMatch) {
			Error(w, http.StatusUnauthorized, err)
			return
		}
		Error(w, http.StatusInternalServerError, err)
		return
	}

	Success(w, http.StatusOK)
}
