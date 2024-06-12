package helpers

import (
	"github.com/golang-jwt/jwt"
)


type IError struct {
	Field string
	Tag   string
	Value string
}

type AuthTokenJwtClaim struct {
	Email  string
	Name   string
	UserId string
	jwt.StandardClaims
}
