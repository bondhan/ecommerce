package model

import (
	"encoding/json"
	ecommerceerror "github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/constants/params"
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	"github.com/go-chi/chi"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
)

type CreatePaymentResp struct {
	PaymentID uint   `json:"paymentId"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreatePaymentReq struct {
	Name string `json:"name"`
}

type Payment struct {
	PaymentID uint   `json:"paymentId"`
	Name      string `json:"name"`
}

type CreatePaymentUpdate struct {
	ID uint `json:"id"`
	CreatePaymentReq
}

func (c CreatePaymentReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
	)
}

func NewPayment(r *http.Request) (CreatePaymentReq, error) {
	var req CreatePaymentReq
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

func UpdatePayment(r *http.Request) (CreatePaymentUpdate, error) {
	IDStr := chi.URLParam(r, "id")
	ID := cast.ToUint(IDStr)
	if ID == 0 {
		return CreatePaymentUpdate{}, ecommerceerror.ErrPaymentNotFound
	}

	req := CreatePaymentUpdate{ID: ID}
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

func GetPaymentID(r *http.Request) (uint, error) {
	IDStr := chi.URLParam(r, "id")
	ID := cast.ToUint(IDStr)
	if ID == 0 {
		return 0, ecommerceerror.ErrPaymentNotFound
	}

	return ID, nil
}

type PaymentPaginatedReq struct {
	Skip  string `json:"skip"`
	Limit string `json:"limit"`
}

func NewPaymentPaginatedReq(r *http.Request) (PaymentPaginated, error) {
	skip := r.URL.Query().Get(params.Skip)
	limit := r.URL.Query().Get(params.Limit)

	categories := PaymentPaginatedReq{
		Skip:  skip,
		Limit: limit,
	}

	err := validation.ValidateStruct(&categories,
		validation.Field(&categories.Limit, is.Digit),
		validation.Field(&categories.Skip, is.Digit),
	)

	var l, s = 0, 0

	if len(categories.Limit) > 0 {
		l, err = strconv.Atoi(categories.Limit)
		if err != nil {
			return PaymentPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	if len(categories.Skip) > 0 {
		s, err = strconv.Atoi(categories.Skip)
		if err != nil {
			return PaymentPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	vp := PaymentPaginated{
		Skip:  s,
		Limit: l,
	}

	return vp, nil
}

type PaymentPaginated struct {
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

type ListResponse struct {
	Categories []Payment      `json:"categories"'`
	Meta       basemodel.Meta `json:"meta"'`
}
