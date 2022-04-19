package usecase

import (
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	"github.com/bondhan/ecommerce/modules/payment/model"
	"github.com/bondhan/ecommerce/modules/payment/query"
	"github.com/sirupsen/logrus"
	"time"
)

type paymentUC struct {
	logger   *logrus.Logger
	paymentQ query.IPaymentQ
}

func NewPaymentUC(logger *logrus.Logger, paymentQ query.IPaymentQ) IPaymentUC {
	return &paymentUC{
		logger:   logger,
		paymentQ: paymentQ,
	}
}
func (c paymentUC) Create(req model.CreatePaymentReq) (model.CreatePaymentResp, error) {
	newPayment, err := c.paymentQ.Insert(req)
	if err != nil {
		return model.CreatePaymentResp{}, err
	}

	nPayment := model.CreatePaymentResp{
		PaymentID: newPayment.ID,
		Name:      newPayment.Name,
		CreatedAt: newPayment.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: newPayment.UpdatedAt.UTC().Format(time.RFC3339),
	}

	return nPayment, nil
}

func (c paymentUC) Update(req model.CreatePaymentUpdate) error {
	err := c.paymentQ.Update(req)
	if err != nil {
		return err
	}

	return nil
}

func (c paymentUC) Delete(id uint) error {
	err := c.paymentQ.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (c paymentUC) List(req model.PaymentPaginated) (model.ListResponse, error) {
	data, count, err := c.paymentQ.List(req)
	if err != nil {
		return model.ListResponse{}, err
	}

	meta := basemodel.Meta{
		Total: count,
		Skip:  req.Skip,
		Limit: req.Limit,
	}

	categories := []model.Payment{}
	for _, v := range data {
		vv := model.Payment{
			Name:      v.Name,
			PaymentID: v.ID,
		}
		categories = append(categories, vv)
	}

	res := model.ListResponse{
		Categories: categories,
		Meta:       meta,
	}

	return res, nil
}
func (c paymentUC) Detail(id uint) (model.Payment, error) {
	data, err := c.paymentQ.Detail(id)
	if err != nil {
		return model.Payment{}, err
	}

	payment := model.Payment{
		Name:      data.Name,
		PaymentID: data.ID,
	}

	return payment, nil
}
