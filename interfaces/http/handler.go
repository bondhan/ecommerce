package handler

import (
	"github.com/bondhan/ecommerce/modules/cashier/presenter"
	"github.com/bondhan/ecommerce/modules/cashier/query"
	"github.com/bondhan/ecommerce/modules/cashier/usecase"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Handler struct {
	cashier presenter.ICashierP
}

func NewHandler(logger *logrus.Logger, db *gorm.DB) *Handler {
	cashierQ := query.NewCashierQ(logger, db)
	cashierUC := usecase.NewCashierUC(logger, cashierQ)
	cashierP := presenter.NewCashierP(cashierUC)

	return &Handler{
		cashier: cashierP,
	}
}
