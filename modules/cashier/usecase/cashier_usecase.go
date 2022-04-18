package usecase

import (
	"github.com/bondhan/ecommerce/modules/cashier/query"
	"github.com/sirupsen/logrus"
)

type cashierUC struct {
	logger   *logrus.Logger
	cashierQ query.ICashierQ
}

func NewCashierUC(logger *logrus.Logger, cashierQ query.ICashierQ) ICashierUC {
	return &cashierUC{
		logger:   logger,
		cashierQ: cashierQ,
	}
}

func (c cashierUC) List()   {}
func (c cashierUC) Detail() {}
func (c cashierUC) Create() {}
func (c cashierUC) Update() {}
func (c cashierUC) Delete() {}
