package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Struct2Map(in interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		out[t.Field(i).Name] = v.Field(i).Interface()
	}
	return out, nil
}

func Slice2String(obj interface{}) (string, error) {
	kind := reflect.TypeOf(obj).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return "", errors.New("obj is not slice or array")
	}
	return strings.Replace(strings.Trim(fmt.Sprint(obj), "[]"), " ", ",", -1), nil
}
