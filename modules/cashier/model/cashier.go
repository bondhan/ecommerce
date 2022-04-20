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

type CreateCashierResp struct {
	Cashier
	PassCode  string `json:"passcode"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreateCashierReq struct {
	Name     string `json:"name"`
	PassCode string `json:"passcode"`
}

type Cashier struct {
	CashierID uint   `json:"cashierId"`
	Name      string `json:"name"`
}

type Passcode struct {
	Passcode string `json:"passcode"`
}

type Token struct {
	Token string `json:"token"`
}
type CreateCashierUpdate struct {
	ID uint
	CreateCashierReq
}

func (c CreateCashierReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.PassCode, validation.Required, is.Digit,
			validation.Length(6, 6)),
	)
}

func NewCashier(r *http.Request) (CreateCashierReq, error) {
	var req CreateCashierReq
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

type CreatePasscodeReq struct {
	PassCode string `json:"passcode"`
}

func (c CreatePasscodeReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.PassCode, validation.Required, is.Digit,
			validation.Length(6, 6)),
	)
}

func NewPasscode(r *http.Request) (CreatePasscodeReq, error) {
	var req CreatePasscodeReq
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

func UpdateCashier(r *http.Request) (CreateCashierUpdate, error) {
	IDStr := chi.URLParam(r, "id")
	ID := cast.ToUint(IDStr)
	if ID == 0 {
		return CreateCashierUpdate{}, ecommerceerror.ErrCashierNotFound
	}

	req := CreateCashierUpdate{ID: ID}
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

func GetCashierID(r *http.Request) (uint, error) {
	IDStr := chi.URLParam(r, "id")
	ID := cast.ToUint(IDStr)
	if ID == 0 {
		return 0, ecommerceerror.ErrCashierNotFound
	}

	return ID, nil
}

type CashierPaginatedReq struct {
	Skip  string `json:"skip"`
	Limit string `json:"limit"`
}

func NewCashierPaginatedReq(r *http.Request) (CashierPaginated, error) {
	skip := r.URL.Query().Get(params.Skip)
	limit := r.URL.Query().Get(params.Limit)

	cashiers := CashierPaginatedReq{
		Skip:  skip,
		Limit: limit,
	}

	err := validation.ValidateStruct(&cashiers,
		validation.Field(&cashiers.Limit, is.Digit),
		validation.Field(&cashiers.Skip, is.Digit),
	)

	var l, s = 0, 0

	if len(cashiers.Limit) > 0 {
		l, err = strconv.Atoi(cashiers.Limit)
		if err != nil {
			return CashierPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	if len(cashiers.Skip) > 0 {
		s, err = strconv.Atoi(cashiers.Skip)
		if err != nil {
			return CashierPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	vp := CashierPaginated{
		Skip:  s,
		Limit: l,
	}

	return vp, nil
}

type CashierPaginated struct {
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

type ListResponse struct {
	Cashiers []Cashier      `json:"cashiers"'`
	Meta     basemodel.Meta `json:"meta"'`
}
