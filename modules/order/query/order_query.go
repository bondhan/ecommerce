package query

import (
	"errors"
	"github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/domain"
	"github.com/bondhan/ecommerce/modules/order/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type orderQ struct {
	logger *logrus.Logger
	gormDB *gorm.DB
}

func NewOrderQ(logger *logrus.Logger, gDB *gorm.DB) IOrderQ {
	return &orderQ{
		logger: logger,
		gormDB: gDB,
	}
}

func (c *orderQ) Insert(req model.OrderTotal, data []model.ProductSubTotal) (domain.Orders, error) {
	newOrder := domain.Orders{
		TotalPrice:    req.TotalPrice,
		TotalPaid:     req.TotalPaid,
		TotalReturn:   req.TotalReturn,
		CashierID:     req.CashiersID,
		PaymentTypeID: req.PaymentTypesID,
		InvoicePDF:    req.ReceiptID,
		Downloaded:    0,
	}

	//example of transactions on golang
	tx := c.gormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(&newOrder).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	for _, v := range data {
		newOrderDetail := domain.OrderDetails{
			Name:             v.Name,
			Price:            v.Price,
			Qty:              v.Qty,
			OrderID:          newOrder.ID,
			ProductID:        v.ProductID,
			TotalNormalPrice: v.TotalNormalPrice,
			TotalFinalPrice:  v.TotalFinalPrice,
		}
		if v.Discount != nil {
			newOrderDetail.DiscountID = v.Discount.ID
		}

		err := tx.Create(&newOrderDetail).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	return newOrder, nil
}

func (c *orderQ) Update(req model.CreateOrderUpdate) error {

	return nil
}

func (c *orderQ) Delete(id uint) error {

	return nil
}

func (c *orderQ) List(req model.OrderPaginated) ([]domain.Orders, int64, error) {
	var orders []domain.Orders
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

	err := c.gormDB.Model(&orders).Count(&count).Error
	if err != nil {
		return orders, count, err
	}

	err = c.gormDB.Limit(limit).Offset(skip).Order("id ASC").Find(&orders).Error
	if err != nil {
		return orders, count, err
	}

	return orders, count, err
}

func (c *orderQ) Detail(id uint) (domain.Orders, error) {
	var order domain.Orders
	err := c.gormDB.Where("id = ?", id).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Orders{}, ecommerceerror.ErrOrderNotFound
		}
		return domain.Orders{}, err
	}
	return order, err
}
