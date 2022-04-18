package presenter

import (
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
		Error(w, http.StatusBadRequest, err)
		return
	}

	JSON(w, http.StatusOK, req)
}
func (c *cashierP) List(w http.ResponseWriter, r *http.Request)   {}
func (c *cashierP) Detail(w http.ResponseWriter, r *http.Request) {}
func (c *cashierP) Update(w http.ResponseWriter, r *http.Request) {}
func (c *cashierP) Delete(w http.ResponseWriter, r *http.Request) {}
