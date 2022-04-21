package usecase

import "github.com/bondhan/ecommerce/modules/category/model"

type ICategoryUC interface {
	Create(req model.CreateCategoryReq) (model.CreateCategoryResp, error)
	Update(req model.CreateCategoryUpdate) error
	Delete(id uint) error
	List(req model.CategoryPaginated) (model.ListResponse, error)
	Detail(id uint) (model.Category, error)
}
