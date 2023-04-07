package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestUpdateProduct(t *testing.T){
	router := setupRouter()

	updateReq := `{"product":{"price":4.55}}`

	req, err := http.NewRequest("PATCH", "/api/v1/products/1", strings.NewReader(updateReq))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if assert.Equal(t, 200, w.Code) {
		assert.JSONEq(t, `{"product":{"id":"1","name":"Apple","price":4.55}}`, w.Body.String())		
	}
}