package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lenuse/mall"
)

func TestCreateAdmin(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/admin/create", "{}")

	router := mall.New()
	router.ServeHTTP(w, req)
}
