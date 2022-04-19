package query

import (
	"github.com/bondhan/ecommerce/domain"
	"github.com/bondhan/ecommerce/modules/product/model"
)

type IProductQ interface {
	Insert(req model.CreateProductReq) (domain.Products, error)
	Update(req model.CreateProductUpdate) error
	Delete(id uint) error
	List(req model.ProductPaginated) ([]model.ProductRes, int64, error)
	Detail(id uint) (model.ProductRes, error)
}
