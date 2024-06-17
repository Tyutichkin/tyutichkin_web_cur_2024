package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Login   string `json:"login"`
	IsAdmin bool   `json:"role"`
	jwt.StandardClaims
}
