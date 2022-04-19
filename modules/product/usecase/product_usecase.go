package usecase

import (
	"fmt"
	"github.com/bondhan/ecommerce/constants/params"
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	"github.com/bondhan/ecommerce/infrastructure"
	modelcategory "github.com/bondhan/ecommerce/modules/category/model"
	"github.com/bondhan/ecommerce/modules/product/model"
	"github.com/bondhan/ecommerce/modules/product/query"
	"github.com/sirupsen/logrus"
	"time"
)

type productUC struct {
	logger   *logrus.Logger
	productQ query.IProductQ
}

func NewProductUC(logger *logrus.Logger, productQ query.IProductQ) IProductUC {
	return &productUC{
		logger:   logger,
		productQ: productQ,
	}
}
func (c productUC) Create(req model.CreateProductReq) (model.CreateProductResp, error) {
	newProduct, err := c.productQ.Insert(req)
	if err != nil {
		return model.CreateProductResp{}, err
	}

	nProduct := model.CreateProductResp{
		ProductID:  newProduct.ID,
		Name:       newProduct.Name,
		CategoryID: newProduct.CategoryID,
		SKU:        fmt.Sprintf("ID%03d", newProduct.ID),
		Image:      newProduct.Image,
		Price:      newProduct.Price,
		Stock:      newProduct.Stock,
		CreatedAt:  newProduct.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:  newProduct.UpdatedAt.UTC().Format(time.RFC3339),
	}

	return nProduct, nil
}

func (c productUC) Update(req model.CreateProductUpdate) error {
	err := c.productQ.Update(req)
	if err != nil {
		return err
	}

	return nil
}

func (c productUC) Delete(id uint) error {
	err := c.productQ.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (c productUC) List(req model.ProductPaginated) (model.ListResponse, error) {
	data, count, err := c.productQ.List(req)
	if err != nil {
		return model.ListResponse{}, err
	}

	meta := basemodel.Meta{
		Total: count,
		Skip:  req.Skip,
		Limit: req.Limit,
	}

	products := []model.ProductDetail{}
	for _, v := range data {
		vv := model.ProductDetail{
			Name:      v.Name,
			ProductID: v.ProductID,
			SKU:       fmt.Sprintf("ID%03d", v.ProductID),
			Image:     v.Image,
			Price:     v.Price,
			Stock:     v.Stock,
			Category: modelcategory.Category{
				CategoryID: v.CategoryID,
				Name:       v.CategoryName,
			},
		}
		if v.DiscountID != nil {
			sf := ""
			if v.Type == params.BuyN {
				sf = fmt.Sprintf("Buy %d only Rp. %s", v.Qty, infrastructure.Dot(v.Result))
			} else if v.Type == params.Percentage {
				sf = fmt.Sprintf("Discout %d%% Rp. %s", v.Result, infrastructure.Dot(v.Price-(v.Result*v.Price/100)))
			}

			vv.Discount = &model.DiscountDetail{
				Qty:             v.Qty,
				Type:            v.Type,
				Result:          v.Result,
				ExpiredAt:       v.ExpiredAt.Format(time.RFC3339),
				ExpiredAtFormat: v.ExpiredAt.Format("02 Jan 2006"),
				StringFormat:    sf,
			}
		}

		products = append(products, vv)
	}

	res := model.ListResponse{
		Products: products,
		Meta:     meta,
	}

	return res, nil
}
func (c productUC) Detail(id uint) (model.ProductDetail, error) {
	v, err := c.productQ.Detail(id)
	if err != nil {
		return model.ProductDetail{}, err
	}

	product := model.ProductDetail{
		Name:      v.Name,
		ProductID: v.ProductID,
		SKU:       fmt.Sprintf("ID%03d", v.ProductID),
		Image:     v.Image,
		Price:     v.Price,
		Stock:     v.Stock,
		Category: modelcategory.Category{
			CategoryID: v.CategoryID,
			Name:       v.CategoryName,
		},
	}
	if v.DiscountID != nil {
		sf := ""
		if v.Type == params.BuyN {
			sf = fmt.Sprintf("Buy %d only Rp. %s", v.Qty, infrastructure.Dot(v.Result))
		} else if v.Type == params.Percentage {
			sf = fmt.Sprintf("Discout %d%% Rp. %s", v.Result, infrastructure.Dot(v.Price-(v.Result*v.Price/100)))
		}

		product.Discount = &model.DiscountDetail{
			ID:              v.DiscountID,
			Qty:             v.Qty,
			Type:            v.Type,
			Result:          v.Result,
			ExpiredAt:       v.ExpiredAt.Format(time.RFC3339),
			ExpiredAtFormat: v.ExpiredAt.Format("02 Jan 2006"),
			StringFormat:    sf,
		}
	}

	return product, nil
}
