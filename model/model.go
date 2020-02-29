package model

import (
	"errors"
	"fmt"
	"mall/conf"
	"os"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Model 接口
type Model interface {
	GetTableName() string
}

var db *gorm.DB
var once sync.Once

func New() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(conf.New().Database, conf.GetHost())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		db.DB().SetMaxIdleConns(conf.New().MaxIdleConns)
		db.DB().SetMaxOpenConns(conf.New().MaxOpenConns)
	})
	return db
}

func BuildJwtToken(authInfo interface{}) (string, error) {
	newTimestamp := time.Now().Unix()
	expiresAt := newTimestamp + 3600
	botBefore := newTimestamp - 1800
	var claims jwt.Claims
	switch authInfo.(type) {
	default:
		return "", errors.New("参数类型错误")
	case AdminClaims:
		claims, _ := authInfo.(AdminClaims)
		claims.ExpiresAt = int64(expiresAt)
		claims.IssuedAt = newTimestamp
		claims.NotBefore = botBefore
		break
	case UserClaims:
		claims, _ := authInfo.(UserClaims)
		claims.ExpiresAt = int64(expiresAt)
		claims.IssuedAt = newTimestamp
		claims.NotBefore = botBefore
		break
	}

	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	keyStr := conf.GetJwtSecret()
	tokenStr, err := tokenStruct.SignedString(keyStr)
	return tokenStr, err
}
