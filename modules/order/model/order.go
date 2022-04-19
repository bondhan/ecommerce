package model

import (
	"encoding/json"
	ecommerceerror "github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/constants/params"
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	"github.com/bondhan/ecommerce/modules/product/model"
	"github.com/go-chi/chi"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
)

type CreateOrderResp struct {
	Order
	PassCode  string `json:"passcode"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreateOrderReq struct {
	Name     string `json:"name"`
	PassCode string `json:"passcode"`
}

type Order struct {
	OrderID uint   `json:"orderId"`
	Name    string `json:"name"`
}

type CreateOrderUpdate struct {
	ID uint
	CreateOrderReq
}

func (c CreateOrderReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.PassCode, validation.Required, is.Digit,
			validation.Length(6, 6)),
	)
}

func NewOrder(r *http.Request) (CreateOrderReq, error) {
	var req CreateOrderReq
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

func UpdateOrder(r *http.Request) (CreateOrderUpdate, error) {
	IDStr := chi.URLParam(r, "id")
	ID := cast.ToUint(IDStr)
	if ID == 0 {
		return CreateOrderUpdate{}, ecommerceerror.ErrOrderNotFound
	}

	req := CreateOrderUpdate{ID: ID}
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

func GetOrderID(r *http.Request) (uint, error) {
	IDStr := chi.URLParam(r, "id")
	ID := cast.ToUint(IDStr)
	if ID == 0 {
		return 0, ecommerceerror.ErrOrderNotFound
	}

	return ID, nil
}

type OrderPaginatedReq struct {
	Skip  string `json:"skip"`
	Limit string `json:"limit"`
}

func NewOrderPaginatedReq(r *http.Request) (OrderPaginated, error) {
	skip := r.URL.Query().Get(params.Skip)
	limit := r.URL.Query().Get(params.Limit)

	orders := OrderPaginatedReq{
		Skip:  skip,
		Limit: limit,
	}

	err := validation.ValidateStruct(&orders,
		validation.Field(&orders.Limit, is.Digit),
		validation.Field(&orders.Skip, is.Digit),
	)

	var l, s = 0, 0

	if len(orders.Limit) > 0 {
		l, err = strconv.Atoi(orders.Limit)
		if err != nil {
			return OrderPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	if len(orders.Skip) > 0 {
		s, err = strconv.Atoi(orders.Skip)
		if err != nil {
			return OrderPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	vp := OrderPaginated{
		Skip:  s,
		Limit: l,
	}

	return vp, nil
}

type OrderPaginated struct {
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

type ListResponse struct {
	Orders []Order        `json:"orders"'`
	Meta   basemodel.Meta `json:"meta"'`
}

type SubTotalReq struct {
	ProductID uint  `json:"productId"'`
	Qty       int64 `json:"qty"'`
}

func (c SubTotalReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ProductID, validation.Required),
		validation.Field(&c.Qty, validation.Required),
	)
}

func NewSubtotalOrder(r *http.Request) ([]SubTotalReq, error) {
	var req []SubTotalReq
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return req, err
	}

	for _, v := range req {
		err := v.Validate()
		if err != nil {
			return req, err
		}
	}

	return req, nil
}

type ProductSubTotal struct {
	ProductID        uint                  `json:"productId"`
	Name             string                `json:"name"`
	Stock            int64                 `json:"stock"`
	Price            int64                 `json:"price"`
	Image            string                `json:"image"`
	Discount         *model.DiscountDetail `json:"discount"`
	Qty              int64                 `json:"qty"`
	TotalNormalPrice int64                 `json:"totalNormalPrice"`
	TotalFinalPrice  int64                 `json:"totalFinalPrice"`
}

type SubTotal struct {
	Subtotal int64             `json:"subtotal"`
	Products []ProductSubTotal `json:"products"`
}
