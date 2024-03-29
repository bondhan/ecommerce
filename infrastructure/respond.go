package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	ecommerceerror "github.com/bondhan/ecommerce/constants/ecommerce_error"
	"net/http"
)

// ErrorResponse is Error response template
type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Code    int         `json:"-"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"-"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *ErrorResponse) String() string {
	return fmt.Sprintf("reason: %s", e.Message)
}

// Respond is response write to ResponseWriter
func Respond(w http.ResponseWriter, code int, src interface{}) {
	var body []byte
	var err error

	switch s := src.(type) {
	case []byte:
		if !json.Valid(s) {
			Error(w, http.StatusInternalServerError, errors.New("Invalid JSON"))
			return
		}
		body = s
	case string:
		body = []byte(s)
	case *ErrorResponse, ErrorResponse:
		// avoid infinite loop
		if body, err = json.Marshal(src); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"reason\":\"failed to parse json\"}"))
			return
		}
	default:
		if body, err = json.Marshal(src); err != nil {
			Error(w, http.StatusInternalServerError, fmt.Errorf("Failed to parse Json: %s", err.Error()))
			return
		}
	}
	if code != http.StatusOK {
		w.WriteHeader(code)
	}
	w.Write(body)
}

// Error is wrapped Respond when error response
func Error(w http.ResponseWriter, code int, err error) {

	c := code

	e := &ErrorResponse{
		Success: false,
		Message: err.Error(),
		Error:   struct{}{},
		Code:    code,
	}

	switch err {
	case ecommerceerror.ErrCashierNotFound, ecommerceerror.ErrCategoryNotFound,
		ecommerceerror.ErrPaymentNotFound, ecommerceerror.ErrProductNotFound, ecommerceerror.ErrOrderNotFound:
		c = http.StatusNotFound
		e.Code = c
	case ecommerceerror.ErrEmptyBody:
		str := `[
        {
            "message": "\"value\" must be an array",
            "path": [],
            "type": "array.base",
            "context": {
                "label": "value",
                "value": {}
            }
			}
		]`
		mm := make([]map[string]interface{}, 1)
		err = json.Unmarshal([]byte(str), &mm)
		if err != nil {
			Error(w, http.StatusInternalServerError, fmt.Errorf("failed to parse Json: %s", err.Error()))
			return
		}
		e.Error = mm
		e.Message = `body ValidationError: "value" must be an array`
	case ecommerceerror.ErrEmptyProduct:
		str := `[{
            "message": "\"products\" is required",
            "path": [
                "products"
            ],
            "type": "any.required",
            "context": {
                "label": "products",
                "key": "products"
            }
        }]`
		mm := make([]map[string]interface{}, 1)
		err = json.Unmarshal([]byte(str), &mm)
		if err != nil {
			Error(w, http.StatusInternalServerError, fmt.Errorf("failed to parse Json: %s", err.Error()))
			return
		}
		e.Error = mm
		e.Message = `body ValidationError:"products" is required`
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, c, e)
}

func JSON(w http.ResponseWriter, code int, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, src)
}

func Success(w http.ResponseWriter, code int) {
	status := Response{
		Success: true,
		Message: "Success",
		Code:    code,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := json.Marshal(status)
	if err != nil {
		Error(w, http.StatusInternalServerError, fmt.Errorf("failed to parse Json: %s", err.Error()))
		return
	}
	if code != http.StatusOK {
		w.WriteHeader(code)
	}
	w.Write(body)
}

func SuccessJSON(w http.ResponseWriter, code int, src interface{}) {
	status := Response{
		Success: true,
		Message: "Success",
		Code:    code,
		Data:    src,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := json.Marshal(status)
	if err != nil {
		Error(w, http.StatusInternalServerError, fmt.Errorf("failed to parse Json: %s", err.Error()))
		return
	}
	if code != http.StatusOK {
		w.WriteHeader(code)
	}
	w.Write(body)
}
