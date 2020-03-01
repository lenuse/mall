package utils

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GetJwtToken 获取jwt签名令牌
func GetJwtToken(id, name, issuer string) (string, error) {
	nowTime := time.Now()
	expiresAt := nowTime.Add(10 * time.Hour)
	beforeAt := nowTime.Add(-3 * time.Minute)
	claims := jwt.StandardClaims{
		Id:        id,
		Subject:   name,
		Issuer:    issuer,
		IssuedAt:  int64(nowTime.Second()),
		ExpiresAt: int64(expiresAt.Second()),
		NotBefore: int64(beforeAt.Second()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString("secret")
}
