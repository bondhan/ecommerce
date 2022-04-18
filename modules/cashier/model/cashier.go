package model

import (
	"encoding/json"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
)

type CreateCashierReq struct {
	Name     string `json:"name"`
	PassCode string `json:"passcode"`
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
