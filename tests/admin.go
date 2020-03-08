package tests

import (
	"github.com/lenuse/mall"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAdmin(t *testing.T)  {
	w := httptest.NewRecorder()
	req ,_ := http.NewRequest("POST", "/v1/admin/create", "{}")

	router := mall.New()
	router.ServeHTTP(w, req)
}