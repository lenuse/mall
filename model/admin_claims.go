package model

import (
	"github.com/dgrijalva/jwt-go"
)

type AdminClaims struct {
	jwt.StandardClaims
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
