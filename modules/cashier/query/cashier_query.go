package query

import (
	"errors"
	"github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/constants/status"
	"github.com/bondhan/ecommerce/domain"
	"github.com/bondhan/ecommerce/modules/cashier/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type cashierQ struct {
	logger *logrus.Logger
	gormDB *gorm.DB
}

func NewCashierQ(logger *logrus.Logger, gDB *gorm.DB) ICashierQ {
	return &cashierQ{
		logger: logger,
		gormDB: gDB,
	}
}

func (c *cashierQ) Insert(req model.CreateCashierReq) (domain.Cashiers, error) {
	newCashier := domain.Cashiers{
		Name:        req.Name,
		Passcode:    req.PassCode,
		LoginStatus: status.LoggedOut,
	}

	err := c.gormDB.Create(&newCashier).Error
	if err != nil {
		return domain.Cashiers{}, err
	}

	return newCashier, nil
}

func (c *cashierQ) Update(req model.CreateCashierUpdate) error {
	var oldCashier domain.Cashiers
	err := c.gormDB.Where("id = ?", req.ID).First(&oldCashier).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ecommerceerror.ErrCashierNotFound
		}
		return err
	}

	res := c.gormDB.Where("id = ? and updated_at = ?", oldCashier.ID, oldCashier.UpdatedAt).Model(&oldCashier).
		Updates(domain.Cashiers{Name: req.Name, Passcode: req.PassCode})
	if err != nil {
		return err
	}

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ecommerceerror.ErrCashierNotFound
	}

	return nil
}

func (c *cashierQ) Delete(id uint) error {
	var cashier domain.Cashiers

	res := c.gormDB.Unscoped().Where("id = ?", id).Model(&cashier).Delete(&cashier)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ecommerceerror.ErrCashierNotFound
	}

	return nil
}

func (c *cashierQ) List(req model.CashierPaginated) ([]domain.Cashiers, int64, error) {
	var cashiers []domain.Cashiers
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

	err := c.gormDB.Model(&cashiers).Count(&count).Error
	if err != nil {
		return cashiers, count, err
	}

	err = c.gormDB.Limit(limit).Offset(skip).Order("id ASC").Find(&cashiers).Error
	if err != nil {
		return cashiers, count, err
	}

	return cashiers, count, err
}

func (c *cashierQ) Detail(id uint) (domain.Cashiers, error) {
	var cashier domain.Cashiers
	err := c.gormDB.Where("id = ?", id).First(&cashier).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Cashiers{}, ecommerceerror.ErrCashierNotFound
		}
		return domain.Cashiers{}, err
	}
	return cashier, err
}
