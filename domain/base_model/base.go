package basemodel

import "github.com/dgrijalva/jwt-go"

type Meta struct {
	Total int64 `json:"total"`
	Skip  int   `json:"skip"`
	Limit int   `json:"limit"`
}

type Claims struct {
	jwt.StandardClaims
}
