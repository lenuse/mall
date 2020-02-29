package model

import (
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
