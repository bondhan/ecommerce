package usecase

import "github.com/bondhan/ecommerce/modules/product/model"

type IProductUC interface {
	Create(req model.CreateProductReq) (model.CreateProductResp, error)
	Update(req model.CreateProductUpdate) error
	Delete(id uint) error
	List(req model.ProductPaginated) (model.ListResponse, error)
	Detail(id uint) (model.ProductDetail, error)
}
