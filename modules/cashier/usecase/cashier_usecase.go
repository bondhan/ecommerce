package usecase

import (
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	"github.com/bondhan/ecommerce/modules/cashier/model"
	"github.com/bondhan/ecommerce/modules/cashier/query"
	"github.com/sirupsen/logrus"
	"time"
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
func (c cashierUC) Create(req model.CreateCashierReq) (model.CreateCashierResp, error) {
	newCashier, err := c.cashierQ.Insert(req)
	if err != nil {
		return model.CreateCashierResp{}, err
	}

	nCashier := model.CreateCashierResp{
		CashierID: newCashier.ID,
		Name:      newCashier.Name,
		PassCode:  newCashier.Passcode,
		CreatedAt: newCashier.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: newCashier.UpdatedAt.UTC().Format(time.RFC3339),
	}

	return nCashier, nil
}

func (c cashierUC) Update(req model.CreateCashierUpdate) error {
	err := c.cashierQ.Update(req)
	if err != nil {
		return err
	}

	return nil
}

func (c cashierUC) Delete(id uint) error {
	err := c.cashierQ.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (c cashierUC) List(req model.CashierPaginated) (model.ListResponse, error) {
	data, count, err := c.cashierQ.List(req)
	if err != nil {
		return model.ListResponse{}, err
	}

	meta := basemodel.Meta{
		Total: count,
		Skip:  req.Skip,
		Limit: req.Limit,
	}

	cashiers := []model.Cashier{}
	for _, v := range data {
		vv := model.Cashier{
			Name:      v.Name,
			CashierID: v.ID,
		}
		cashiers = append(cashiers, vv)
	}

	res := model.ListResponse{
		Cashiers: cashiers,
		Meta:     meta,
	}

	return res, nil
}
func (c cashierUC) Detail(id uint) (model.Cashier, error) {
	data, err := c.cashierQ.Detail(id)
	if err != nil {
		return model.Cashier{}, err
	}

	cashier := model.Cashier{
		Name:      data.Name,
		CashierID: data.ID,
	}

	return cashier, nil
}
