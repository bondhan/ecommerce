package middleware

import (
	"context"
	ecommerceerror "github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/constants/params"
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	. "github.com/bondhan/ecommerce/infrastructure"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/middleware"
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	dbLog *gorm.DB
	mutex sync.Mutex
)

// get the requestId from header 'X-Request-Id' if not then generate one from UUID
func RequestUUID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(middleware.RequestIDHeader)
		if requestID == "" {
			requestID = uuid.New().String() //use google uuid
		}

		ctx = context.WithValue(ctx, middleware.RequestIDKey, requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TimeoutHandler(timeout time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if next == nil {
				next = http.DefaultServeMux
			}
			timeoutHandler := http.TimeoutHandler(next, timeout, "Request Timeout")
			timeoutHandler.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}

func JWTValidator(jwtKey string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		claims := &basemodel.Claims{}
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				Error(w, http.StatusUnauthorized, ecommerceerror.ErrUnauthorized)
				return
			}
			Error(w, http.StatusBadRequest, ecommerceerror.ErrUnauthorized)
			return
		}
		if !tkn.Valid {
			Error(w, http.StatusUnauthorized, ecommerceerror.ErrUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, params.JWTClaims, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
