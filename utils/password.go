package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"math"

	"golang.org/x/crypto/bcrypt"
)

// M5Encrypt 创建md5字符串
func M5Encrypt(str string) string {
	h := md5.New()
	tag := []byte(str)
	h.Write(tag)
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// GetPseudoRandomString 创建定长随机字符串
func GetPseudoRandomString(num uint) (string, error) {
	byteLen := math.Ceil(float64(num / 2))
	bytes := make([]byte, int(byteLen))
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	randStr := hex.EncodeToString(bytes)
	return randStr[0:num], nil
}

// PasswordHash 通过bcrypt算法生成hash
func PasswordHash(password []byte) (string, error) {

	hashBytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hashBytes), err
}

// PasswordVerify 通过bcrypt算法验证hash
func PasswordVerify(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
