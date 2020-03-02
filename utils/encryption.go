package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GetBcryptHash 生成 bcrypt的hash字符串
func GetBcryptHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// VerifyBcryptHash 验证hash与原文
func VerifyBcryptHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
