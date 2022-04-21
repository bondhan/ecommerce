package usecase

import (
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	"github.com/bondhan/ecommerce/modules/category/model"
	"github.com/bondhan/ecommerce/modules/category/query"
	"github.com/sirupsen/logrus"
	"time"
)

type categoryUC struct {
	logger    *logrus.Logger
	categoryQ query.ICategoryQ
}

func NewCategoryUC(logger *logrus.Logger, categoryQ query.ICategoryQ) ICategoryUC {
	return &categoryUC{
		logger:    logger,
		categoryQ: categoryQ,
	}
}
func (c categoryUC) Create(req model.CreateCategoryReq) (model.CreateCategoryResp, error) {
	newCategory, err := c.categoryQ.Insert(req)
	if err != nil {
		return model.CreateCategoryResp{}, err
	}

	nCategory := model.CreateCategoryResp{
		Category: model.Category{
			CategoryID: newCategory.ID,
			Name:       newCategory.Name,
		},
		CreatedAt: newCategory.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: newCategory.UpdatedAt.UTC().Format(time.RFC3339),
	}

	return nCategory, nil
}

func (c categoryUC) Update(req model.CreateCategoryUpdate) error {
	err := c.categoryQ.Update(req)
	if err != nil {
		return err
	}

	return nil
}

func (c categoryUC) Delete(id uint) error {
	err := c.categoryQ.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (c categoryUC) List(req model.CategoryPaginated) (model.ListResponse, error) {
	data, count, err := c.categoryQ.List(req)
	if err != nil {
		return model.ListResponse{}, err
	}

	meta := basemodel.Meta{
		Total: count,
		Skip:  req.Skip,
		Limit: req.Limit,
	}

	categories := []model.Category{}
	for _, v := range data {
		vv := model.Category{
			Name:       v.Name,
			CategoryID: v.ID,
		}
		categories = append(categories, vv)
	}

	res := model.ListResponse{
		Categories: categories,
		Meta:       meta,
	}

	return res, nil
}
func (c categoryUC) Detail(id uint) (model.Category, error) {
	data, err := c.categoryQ.Detail(id)
	if err != nil {
		return model.Category{}, err
	}

	category := model.Category{
		Name:       data.Name,
		CategoryID: data.ID,
	}

	return category, nil
}
