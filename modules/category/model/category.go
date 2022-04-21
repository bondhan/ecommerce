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

type CreateCategoryResp struct {
	Category
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreateCategoryReq struct {
	Name string `json:"name"`
}

type Category struct {
	CategoryID uint   `json:"categoryId"`
	Name       string `json:"name"`
}

type CreateCategoryUpdate struct {
	ID uint
	CreateCategoryReq
}

func (c CreateCategoryReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
	)
}

func NewCategory(r *http.Request) (CreateCategoryReq, error) {
	var req CreateCategoryReq
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

func UpdateCategory(r *http.Request) (CreateCategoryUpdate, error) {
	IDStr := chi.URLParam(r, "id")
	ID := cast.ToUint(IDStr)
	if ID == 0 {
		return CreateCategoryUpdate{}, ecommerceerror.ErrCategoryNotFound
	}

	req := CreateCategoryUpdate{ID: ID}
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

func GetCategoryID(r *http.Request) (uint, error) {
	IDStr := chi.URLParam(r, "id")
	ID := cast.ToUint(IDStr)
	if ID == 0 {
		return 0, ecommerceerror.ErrCategoryNotFound
	}

	return ID, nil
}

type CategoryPaginatedReq struct {
	Skip  string `json:"skip"`
	Limit string `json:"limit"`
}

func NewCategoryPaginatedReq(r *http.Request) (CategoryPaginated, error) {
	skip := r.URL.Query().Get(params.Skip)
	limit := r.URL.Query().Get(params.Limit)

	categories := CategoryPaginatedReq{
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
			return CategoryPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	if len(categories.Skip) > 0 {
		s, err = strconv.Atoi(categories.Skip)
		if err != nil {
			return CategoryPaginated{}, ecommerceerror.ErrInvalidParameters
		}
	}

	vp := CategoryPaginated{
		Skip:  s,
		Limit: l,
	}

	return vp, nil
}

type CategoryPaginated struct {
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

type ListResponse struct {
	Categories []Category     `json:"categories"`
	Meta       basemodel.Meta `json:"meta"`
}
