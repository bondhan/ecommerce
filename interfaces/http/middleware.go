package handler

import (
	"context"
	ecommerceerror "github.com/bondhan/ecommerce/constants/ecommerce_error"
	"github.com/bondhan/ecommerce/constants/params"
	basemodel "github.com/bondhan/ecommerce/domain/base_model"
	. "github.com/bondhan/ecommerce/infrastructure"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

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
