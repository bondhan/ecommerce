package query

import (
	"github.com/bondhan/ecommerce/domain"
	"github.com/bondhan/ecommerce/modules/category/model"
)

type ICategoryQ interface {
	Insert(req model.CreateCategoryReq) (domain.Category, error)
	Update(req model.CreateCategoryUpdate) error
	Delete(id uint) error
	List(req model.CategoryPaginated) ([]domain.Category, int64, error)
	Detail(id uint) (domain.Category, error)
}
