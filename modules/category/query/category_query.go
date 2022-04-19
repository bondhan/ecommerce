package query

import (
	"errors"
	"github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/domain"
	"github.com/bondhan/ecommerce/modules/category/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type categoryQ struct {
	logger *logrus.Logger
	gormDB *gorm.DB
}

func NewCategoryQ(logger *logrus.Logger, gDB *gorm.DB) ICategoryQ {
	return &categoryQ{
		logger: logger,
		gormDB: gDB,
	}
}

func (c *categoryQ) Insert(req model.CreateCategoryReq) (domain.Category, error) {
	newCategory := domain.Category{
		Name: req.Name,
	}

	err := c.gormDB.Create(&newCategory).Error
	if err != nil {
		return domain.Category{}, err
	}

	return newCategory, nil
}

func (c *categoryQ) Update(req model.CreateCategoryUpdate) error {
	var oldCategory domain.Category
	err := c.gormDB.Where("id = ?", req.ID).First(&oldCategory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ecommerceerror.ErrCategoryNotFound
		}
		return err
	}

	res := c.gormDB.Where("id = ? and updated_at = ?", oldCategory.ID, oldCategory.UpdatedAt).Model(&oldCategory).
		Updates(domain.Category{Name: req.Name})
	if err != nil {
		return err
	}

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ecommerceerror.ErrCategoryNotFound
	}

	return nil
}

func (c *categoryQ) Delete(id uint) error {
	var category domain.Category

	res := c.gormDB.Unscoped().Where("id = ?", id).Model(&category).Delete(&category)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ecommerceerror.ErrCategoryNotFound
	}

	return nil
}

func (c *categoryQ) List(req model.CategoryPaginated) ([]domain.Category, int64, error) {
	var categories []domain.Category
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

	err := c.gormDB.Model(&categories).Count(&count).Error
	if err != nil {
		return categories, count, err
	}

	err = c.gormDB.Limit(limit).Offset(skip).Order("id ASC").Find(&categories).Error
	if err != nil {
		return categories, count, err
	}

	return categories, count, err
}

func (c *categoryQ) Detail(id uint) (domain.Category, error) {
	var category domain.Category
	err := c.gormDB.Where("id = ?", id).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Category{}, ecommerceerror.ErrCategoryNotFound
		}
		return domain.Category{}, err
	}
	return category, err
}
