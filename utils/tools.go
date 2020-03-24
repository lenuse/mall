package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Struct2Map(obj interface{}) (map[string]interface{}, error) {
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Struct {
		return nil, errors.New("obj is not struct")
	}
	objValue := reflect.ValueOf(obj)
	data := make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		data[objType.Field(i).Name] = objValue.Field(i).Interface()
	}
	return data, nil
}

func Slice2String(obj interface{}) (string, error) {
	kind := reflect.TypeOf(obj).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return "", errors.New("obj is not slice or array")
	}
	return strings.Replace(strings.Trim(fmt.Sprint(obj), "[]"), " ", ",", -1), nil
}
