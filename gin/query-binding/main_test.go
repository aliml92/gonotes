package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestFetchFlights(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/api/v1/flights?from=SFO&to=LAX&offset=3&limit=5", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	expected := `[
		{"flight_id":"flight3","departure_airport_code":"SFO","arrival_airport_code":"LAX"},
		{"flight_id":"flight4","departure_airport_code":"SFO","arrival_airport_code":"LAX"},
		{"flight_id":"flight5","departure_airport_code":"SFO","arrival_airport_code":"LAX"},
		{"flight_id":"flight6","departure_airport_code":"SFO","arrival_airport_code":"LAX"},
		{"flight_id":"flight7","departure_airport_code":"SFO","arrival_airport_code":"LAX"}
	]`

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, expected, w.Body.String())
}