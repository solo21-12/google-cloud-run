package models

import "github.com/dgrijalva/jwt-go"

type JWTCustome struct {
	Database string `json:"database_name"`
	Expires  int64  `json:"expires"`
	jwt.StandardClaims
}
