package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	// this is like a database
	store := NewInMemoryPlayerStore()
	
	// main server
	server := PlayerServer{store}

	// user
	player := "Pepper"

	// user making POST request to record his win x3
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	// at this point, winningScore for the user "Pepper" must be 3
	
	// now this time, the user is making GET request to see his score
	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

