package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lenuse/mall/repository"
	"github.com/lenuse/mall/utils"
	"github.com/stretchr/testify/assert"

	"github.com/lenuse/mall"
)

func TestCreateAdmin(t *testing.T) {
	repository.Init()
	defer repository.Close()
	w := httptest.NewRecorder()
	jsonStream := strings.NewReader(`{"username":"admin","password":"a1310","repeat_input":"a1310","nickname":"lenu","email":"283689464@qq.com","note":""}`)
	req, _ := http.NewRequest("POST", "/v1/admin/create", jsonStream)

	router := mall.New()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	bodyMap := make(map[string]interface{})
	respStream := w.Body.Bytes()
	err := json.Unmarshal(respStream, bodyMap)
	if err != nil {
		assert.Errorf(t, err, "转化json失败：%s", err.Error())
	}
	state, ok := bodyMap["state"]
	if !ok {
		assert.Error(t, errors.New("没有state"))
	}
	stateCode, _ := state.(int)
	if stateCode != utils.Success.Int() {
		assert.Errorf(t, errors.New("state 不为成功"), "state is %d", stateCode)
	}
}

func TestAdminLogin(t *testing.T) {
	repository.Init()
	defer repository.Close()
	w := httptest.NewRecorder()
	jsonStream := strings.NewReader(`{"username":"admin","password":"a1310"}`)
	req, _ := http.NewRequest("POST", "/v1/admin/login", jsonStream)
	router := mall.New()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	bodyMap := make(map[string]interface{})
	respStream := w.Body.Bytes()
	err := json.Unmarshal(respStream, bodyMap)
	if err != nil {
		assert.Errorf(t, err, "转化json失败：%s", err.Error())
	}
	state, ok := bodyMap["state"]
	if !ok {
		assert.Error(t, errors.New("没有state"))
	}
	stateCode, _ := state.(int)
	if stateCode != utils.Success.Int() {
		assert.Errorf(t, errors.New("state 不为成功"), "state is %d", stateCode)
	}
}

func BenchmarkAdminLogin(b *testing.B) {
	_ = repository.Open()
	defer repository.Close()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		jsonStream := strings.NewReader(`{"username":"admin","password":"a1310"}`)
		req, _ := http.NewRequest("POST", "/v1/admin/login", jsonStream)
		router := mall.New()
		router.ServeHTTP(w, req)
		assert.Equal(b, 200, w.Code)
		bodyMap := make(map[string]interface{})
		respStream := w.Body.Bytes()
		err := json.Unmarshal(respStream, bodyMap)
		if err != nil {
			assert.Errorf(b, err, "转化json失败：%s", err.Error())
		}
		state, ok := bodyMap["state"]
		if !ok {
			assert.Error(b, errors.New("没有state"))
		}
		stateCode, _ := state.(int)
		if stateCode != utils.Success.Int() {
			assert.Errorf(b, errors.New("state 不为成功"), "state is %d", stateCode)
		}
	}

}
