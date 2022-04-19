package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ErrorResponse is Error response template
type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"-"`
}

type Response struct {
	Status  bool        `json:"status"`
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
	w.WriteHeader(code)
	w.Write(body)
}

// Error is wrapped Respond when error response
func Error(w http.ResponseWriter, code int, err error) {
	e := &ErrorResponse{
		Status:  false,
		Message: err.Error(),
		Code:    code,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, e)
}

func JSON(w http.ResponseWriter, code int, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, src)
}

func Success(w http.ResponseWriter, code int) {
	status := Response{
		Status:  true,
		Message: "Success",
		Code:    code,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := json.Marshal(status)
	if err != nil {
		Error(w, http.StatusInternalServerError, fmt.Errorf("failed to parse Json: %s", err.Error()))
		return
	}
	w.WriteHeader(code)
	w.Write(body)
}

func SuccessJSON(w http.ResponseWriter, code int, src interface{}) {
	status := Response{
		Status:  true,
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
	w.WriteHeader(code)
	w.Write(body)
}
