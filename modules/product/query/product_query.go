package query

import (
	"errors"
	"fmt"
	"github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/domain"
	"github.com/bondhan/ecommerce/modules/product/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type productQ struct {
	logger *logrus.Logger
	gormDB *gorm.DB
}

func NewProductQ(logger *logrus.Logger, gDB *gorm.DB) IProductQ {
	return &productQ{
		logger: logger,
		gormDB: gDB,
	}
}

func (c *productQ) Insert(req model.CreateProductReq) (domain.Products, error) {
	newProduct := domain.Products{
		Name:       req.Name,
		Image:      req.Image,
		Price:      req.Price,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
	}

	var disc *domain.Discount
	if req.Discount != nil {
		t := time.Unix(req.Discount.ExpiredAt, 0)
		exp := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)

		disc = &domain.Discount{
			Qty:       req.Discount.Qty,
			Type:      req.Discount.Type,
			Result:    req.Discount.Result,
			ExpiredAt: exp,
		}
	}

	//example of transactions on golang
	tx := c.gormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if disc != nil {
		err := tx.Create(disc).Error
		if err != nil {
			tx.Rollback()
			return domain.Products{}, err
		}

		newProduct.DiscountID = &disc.ID
	}

	err := tx.Create(&newProduct).Error
	if err != nil {
		tx.Rollback()
		return domain.Products{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return domain.Products{}, err
	}

	return newProduct, nil
}

func (c *productQ) Update(req model.CreateProductUpdate) error {
	var oldProduct domain.Products
	err := c.gormDB.Where("id = ?", req.ID).First(&oldProduct).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ecommerceerror.ErrProductNotFound
		}
		return err
	}

	newProduct := domain.Products{
		Name:       req.Name,
		Image:      req.Image,
		CategoryID: req.CategoryID,
		Price:      req.Price,
		Stock:      req.Stock,
	}

	res := c.gormDB.Where("id = ? and updated_at = ?", oldProduct.ID, oldProduct.UpdatedAt).Model(&oldProduct).
		Updates(newProduct)
	if err != nil {
		return err
	}

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ecommerceerror.ErrProductNotFound
	}

	return nil
}

func (c *productQ) Delete(id uint) error {
	var product domain.Products
	var discount domain.Discount
	var oldProduct domain.Products

	err := c.gormDB.Where("id = ?", id).First(&oldProduct).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ecommerceerror.ErrProductNotFound
		}
		return err
	}

	//example of transactions on golang
	tx := c.gormDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	res := tx.Unscoped().Where("id = ?", id).Model(&product).Delete(&product)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if oldProduct.DiscountID != nil {
		res = tx.Unscoped().Where("id = ?", oldProduct.DiscountID).Model(&discount).Delete(&discount)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if res.RowsAffected == 0 {
		return ecommerceerror.ErrProductNotFound
	}

	return nil
}

func (c *productQ) List(req model.ProductPaginated) ([]model.ProductRes, int64, error) {
	products := []model.ProductRes{}
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
			products p
		left join discounts d on
			d.id = p.discount_id
		inner join categories c on
			c.id = p.category_id
		where
			p.deleted_at is null
	`

	wquery := "1=1"
	if req.CategoryID != nil {
		wquery = fmt.Sprintf("%s and p.category_id = %d", wquery, *req.CategoryID)
	}

	if req.Query != nil {
		wquery = fmt.Sprintf("%s and p.name like '%%%s%%'", wquery, *req.Query)
	}

	err := c.gormDB.Raw(fmt.Sprintf("%s and %s", queryC, wquery)).Model(&domain.Products{}).Count(&count).Error
	if err != nil {
		return products, count, err
	}

	query := `
		select
			p.id,
			p.name,
			p.stock,
			p.price,
			p.image,
			p.category_id,
			c.name as category_name,
			p.discount_id,
			d.type,
			d.qty,
			d.result,
			d.expired_at
		from
			products p
		left join discounts d on
			d.id = p.discount_id
		inner join categories c on
			c.id = p.category_id
		where
			p.deleted_at is null
	`

	qq := fmt.Sprintf("%s and %s LIMIT %d OFFSET %d", query, wquery, limit, skip)

	err = c.gormDB.Raw(qq).
		Limit(limit).Offset(skip).Order("id ASC").Scan(&products).Error
	if err != nil {
		return products, count, err
	}

	return products, count, err
}

func (c *productQ) Detail(id uint) (model.ProductRes, error) {
	product := model.ProductRes{}
	query := `
		select
			p.id,
			p.name,
			p.stock,
			p.price,
			p.image,
			p.category_id,
			c.name as category_name,
			p.discount_id,
			d.type,
			d.qty,
			d.result,
			d.expired_at
		from
			products p
		left join discounts d on
			d.id = p.discount_id
		inner join categories c on
			c.id = p.category_id
		where
			p.deleted_at is null
			and p.id = %d
	`

	q := fmt.Sprintf(query, id)

	err := c.gormDB.Raw(q).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ProductRes{}, ecommerceerror.ErrProductNotFound
		}
		return model.ProductRes{}, err
	}
	return product, err
}
