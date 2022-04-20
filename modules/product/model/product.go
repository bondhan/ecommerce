package model

import (
	"encoding/json"
	ecommerceerror "github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/constants/params"
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	modelcategory "github.com/bondhan/ecommerce/modules/category/model"

	"github.com/go-chi/chi"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
	"time"
)

type CreateProductResp struct {
	ProductID  uint   `json:"productId"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	SKU        string `json:"sku"`
	Price      int64  `json:"price"`
	Stock      int64  `json:"stock"`
	CategoryID uint   `json:"categoriesId"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type Discount struct {
	Qty       int64  `json:"qty"`
	Type      string `json:"type"`
	Result    int64  `json:"result"`
	ExpiredAt int64  `json:"expiredAt"`
}

type CreateProductReq struct {
	Name       string    `json:"name"`
	Image      string    `json:"image"`
	Price      int64     `json:"price"`
	Stock      int64     `json:"stock"`
	CategoryID uint      `json:"categoryId"`
	Discount   *Discount `json:"discount"`
}

type DiscountDetail struct {
	ID              *uint  `json:"discountId"`
	Qty             int64  `json:"qty"`
	Type            string `json:"type"`
	Result          int64  `json:"result"`
	ExpiredAt       string `json:"expiredAt"`
	ExpiredAtFormat string `json:"expiredAtFormat"`
	StringFormat    string `json:"stringFormat"`
}

type ProductDetail struct {
	ProductID uint                   `json:"productId"`
	Name      string                 `json:"name"`
	SKU       string                 `json:"sku"`
	Image     string                 `json:"image"`
	Price     int64                  `json:"price"`
	Stock     int64                  `json:"stock"`
	Discount  *DiscountDetail        `json:"discount"`
	Category  modelcategory.Category `json:"category"`
}

type CreateProductUpdate struct {
	ID uint
	CreateProductReq
}

func (c CreateProductReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
	)
}

func NewProduct(r *http.Request) (CreateProductReq, error) {
	var req CreateProductReq
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return req, err
	}

	err := req.Validate()
	if err != nil {
		return req, err
	}

	return req, nil
}

func UpdateProduct(r *http.Request) (CreateProductUpdate, error) {
	IDStr := chi.URLParam(r, "id")
	ID := cast.ToUint(IDStr)
	if ID == 0 {
		return CreateProductUpdate{}, ecommerceerror.ErrProductNotFound
	}

	req := CreateProductUpdate{ID: ID}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return req, err
	}

	err := req.Validate()
	if err != nil {
		return req, err
	}

	return req, nil
}

func GetProductID(r *http.Request) (uint, error) {
	IDStr := chi.URLParam(r, "id")
	ID := cast.ToUint(IDStr)
	if ID == 0 {
		return 0, ecommerceerror.ErrProductNotFound
	}

	return ID, nil
}

type ProductPaginatedReq struct {
	Skip       string  `json:"skip"`
	Limit      string  `json:"limit"`
	CategoryID *string `json:"categoryId"`
	Query      *string `json:"query"`
}

func NewProductPaginatedReq(r *http.Request) (ProductPaginated, error) {
	skip := r.URL.Query().Get(params.Skip)
	limit := r.URL.Query().Get(params.Limit)
	categoryID := r.URL.Query().Get(params.CategoryId)
	query := r.URL.Query().Get(params.Query)

	products := ProductPaginatedReq{
		Skip:       skip,
		Limit:      limit,
		CategoryID: &categoryID,
		Query:      &query,
	}

	err := validation.ValidateStruct(&products,
		validation.Field(&products.Limit, is.Digit),
		validation.Field(&products.Skip, is.Digit),
		validation.Field(&products.CategoryID, is.Digit),
	)

	var l, s = 0, 0

	if len(products.Limit) > 0 {
		l, err = strconv.Atoi(products.Limit)
		if err != nil {
			return ProductPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	if len(products.Skip) > 0 {
		s, err = strconv.Atoi(products.Skip)
		if err != nil {
			return ProductPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	var c *int
	if products.CategoryID != nil && len(*products.CategoryID) > 0 {
		cc, err := strconv.Atoi(*products.CategoryID)
		c = &cc
		if err != nil {
			return ProductPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	var q *string
	if len(query) > 0 {
		q = &query
	}

	vp := ProductPaginated{
		Skip:       s,
		Limit:      l,
		CategoryID: c,
		Query:      q,
	}

	return vp, nil
}

type ProductPaginated struct {
	Skip       int     `json:"skip"`
	Limit      int     `json:"limit"`
	CategoryID *int    `json:"categoryId"`
	Query      *string `json:"query"`
}

type ListResponse struct {
	Products []ProductDetail `json:"products"`
	Meta     basemodel.Meta  `json:"meta"`
}

type ProductRes struct {
	ProductID    uint      `gorm:"column:id"`
	Name         string    `gorm:"column:name"`
	Image        string    `gorm:"column:image"`
	Stock        int64     `gorm:"column:stock"`
	Price        int64     `gorm:"column:price"`
	CategoryID   uint      `gorm:"column:category_id"`
	CategoryName string    `gorm:"column:category_name"`
	DiscountID   *uint     `gorm:"column:discount_id"`
	Qty          int64     `gorm:"column:qty"`
	Type         string    `gorm:"column:type"`
	Result       int64     `gorm:"column:result"`
	ExpiredAt    time.Time `gorm:"column:expired_at"`
}
