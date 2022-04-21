package query

import (
	"errors"
	"fmt"
	"github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/domain"
	"github.com/bondhan/ecommerce/modules/order/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
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
		CashierID:     cast.ToUint(req.CashiersID),
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
		// get product stock
		product := domain.Products{}
		err = tx.Where("id = ? and deleted_at is null", v.ProductID).First(&product).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}

		// sub stock
		remainStock := product.Stock - v.Qty
		if remainStock < 0 {
			tx.Rollback()
			return domain.Orders{}, ecommerceerror.ErrOutOfStock
		}

		// update stock
		res := tx.Where("id = ? and updated_at = ?", product.ID, product.UpdatedAt).Model(&product).
			Updates(map[string]interface{}{"stock": remainStock})
		if res.Error != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}

		if res.RowsAffected == 0 {
			tx.Rollback()
			return domain.Orders{}, ecommerceerror.ErrStockChange
		}

		err = tx.Create(&newOrderDetail).Error
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

func (c *orderQ) List(req model.OrderPaginated) ([]model.OrderDetailDB, int64, error) {
	orders := []model.OrderDetailDB{}
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
		//skip = (skip - 1) * limit
	}

	queryC := `
		select
			count(*)
		from
			orders o
		inner join payment_types pt on
			pt.id = o.payment_type_id
		inner join cashiers c on
			c.id = o.cashier_id
		where
			o.deleted_at is null
			and pt.deleted_at is NULL 
			and c.deleted_at is null
	`

	err := c.gormDB.Raw(queryC).Model(&domain.Products{}).Count(&count).Error
	if err != nil {
		return orders, count, err
	}

	query := `
		select
			o.id,
			o.cashier_id,
			o.payment_type_id,
			o.total_price,
			o.total_paid,
			o.total_return,
			o.updated_at,
			o.created_at,
			c.id as cashier_id,
			c.name as cashier_name,
			pt.id as payment_type_id, 
			pt.name as payment_name,
			pt.type as payment_type,
			pt.logo as payment_logo	
		from
			orders o
		inner join payment_types pt on
			pt.id = o.payment_type_id
		inner join cashiers c on
			c.id = o.cashier_id
		where
			o.deleted_at is null
			and pt.deleted_at is NULL 
			and c.deleted_at is null
	`

	qq := fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, skip)

	err = c.gormDB.Raw(qq).
		Limit(limit).Offset(skip).Order("id ASC").Scan(&orders).Error
	if err != nil {
		return orders, count, err
	}

	return orders, count, err
}

func (c *orderQ) DetailOrdersById(id uint) (*model.OrderDetailDB, error) {
	var order *model.OrderDetailDB

	query := `
		select
			o.id,
			o.cashier_id,
			o.payment_type_id,
			o.total_price,
			o.total_paid,
			o.total_return,
			o.updated_at,
			o.created_at,
			c.id as cashier_id,
			c.name as cashier_name,
			pt.id as payment_type_id, 
			pt.name as payment_name,
			pt.type as payment_type,
			pt.logo as payment_logo	
		from
			orders o
		inner join payment_types pt on
			pt.id = o.payment_type_id
		inner join cashiers c on
			c.id = o.cashier_id
		where
			o.deleted_at is null
			and pt.deleted_at is NULL 
			and c.deleted_at is null
			and o.id = ?
	`

	err := c.gormDB.Raw(query, id).Scan(&order).Error
	if err != nil {
		return order, err
	}

	if order == nil {
		return order, ecommerceerror.ErrOrderNotFound
	}

	return order, err
}

func (c *orderQ) ListOrderDetailsByOrderId(id uint) ([]domain.OrderDetails, error) {
	var orderDetails []domain.OrderDetails
	err := c.gormDB.Where("order_id = ?", id).First(&orderDetails).Error
	if err != nil {
		return []domain.OrderDetails{}, err
	}
	return orderDetails, err
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
	return order, nil
}

func (c *orderQ) SetDownload(id uint, status int) ([]byte, error) {

	var order domain.Orders
	err := c.gormDB.Where("id = ?", id).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecommerceerror.ErrOrderNotFound
		}
		return nil, err
	}

	tx := c.gormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	res := c.gormDB.Where("id = ? and updated_at = ?", order.ID, order.UpdatedAt).Model(&order).
		Updates(map[string]interface{}{"downloaded": status})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, ecommerceerror.ErrOrderNotFound
		}
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, ecommerceerror.ErrOrderNotFound
	}

	return []byte(fmt.Sprintf("%+v", order)), nil
}
