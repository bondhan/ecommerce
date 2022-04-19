package query

import (
	"errors"
	"github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/domain"
	"github.com/bondhan/ecommerce/modules/payment/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type paymentQ struct {
	logger *logrus.Logger
	gormDB *gorm.DB
}

func NewPaymentQ(logger *logrus.Logger, gDB *gorm.DB) IPaymentQ {
	return &paymentQ{
		logger: logger,
		gormDB: gDB,
	}
}

func (c *paymentQ) Insert(req model.CreatePaymentReq) (domain.Payments, error) {
	newPayment := domain.Payments{
		Name: req.Name,
		Type: req.Type,
		Logo: req.Logo,
	}

	err := c.gormDB.Create(&newPayment).Error
	if err != nil {
		return domain.Payments{}, err
	}

	return newPayment, nil
}

func (c *paymentQ) Update(req model.CreatePaymentUpdate) error {
	var oldPayment domain.Payments
	err := c.gormDB.Where("id = ?", req.ID).First(&oldPayment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ecommerceerror.ErrPaymentNotFound
		}
		return err
	}

	res := c.gormDB.Where("id = ? and updated_at = ?", oldPayment.ID, oldPayment.UpdatedAt).Model(&oldPayment).
		Updates(domain.Payments{Name: req.Name, Type: req.Type, Logo: req.Logo})
	if err != nil {
		return err
	}

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ecommerceerror.ErrPaymentNotFound
	}

	return nil
}

func (c *paymentQ) Delete(id uint) error {
	var payment domain.Payments

	res := c.gormDB.Unscoped().Where("id = ?", id).Model(&payment).Delete(&payment)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ecommerceerror.ErrPaymentNotFound
	}

	return nil
}

func (c *paymentQ) List(req model.PaymentPaginated) ([]domain.Payments, int64, error) {
	var payments []domain.Payments
	var limit, skip = 0, 0
	var count int64

	limit = req.Limit
	if req.Limit < 1 {
		limit = 10
	}

	skip = req.Skip
	if req.Skip < 1 {
		skip = 0
	} else {
		skip = (skip - 1) * limit
	}

	err := c.gormDB.Model(&payments).Count(&count).Error
	if err != nil {
		return payments, count, err
	}

	err = c.gormDB.Limit(limit).Offset(skip).Order("id ASC").Find(&payments).Error
	if err != nil {
		return payments, count, err
	}

	return payments, count, err
}

func (c *paymentQ) Detail(id uint) (domain.Payments, error) {
	var payment domain.Payments
	err := c.gormDB.Where("id = ?", id).First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Payments{}, ecommerceerror.ErrPaymentNotFound
		}
		return domain.Payments{}, err
	}
	return payment, err
}
