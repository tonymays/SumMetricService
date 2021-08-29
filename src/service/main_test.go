package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ta App

func setup() {
	ta = App{}
	ta.Init()

}

func executeRequest(ta App, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ta.Server.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetMetric(t *testing.T) {
	setup()
	req, _ := http.NewRequest("GET", "/metrics/active_visitors", nil)
	req.Header.Add("Content-Type", "application/json")
	testResponse := executeRequest(ta, req)
	checkResponseCode(t, http.StatusOK, testResponse.Code)
}


/*

var ta App

// Setup Tests
func setup() {
	ta = App{}
	ta.Init()
}

// Execute test an http request // TODO
func executeRequest(ta App, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ta.Server.Router.ServeHTTP(rr, req)
	return rr
}

// Cleanup DB CACHE after a test
func clean(ta App) {
	// TODO clean DB, make new TestDataCache
}

// Check response code returned from a test http request
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// TestPostMetric Test
func TestGetMetric(t *testing.T) {
	// Test Setup
	setup()

	// Create User Test Request
	payload := []byte(`{"Key":"EXAMPLE","Value":"EXAMPLE"}`)
	req, _ := http.NewRequest("GET", "/metrics/active_visitors", bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")
	testResponse := executeRequest(ta, req)
	// Clean test database and check test response
	clean(ta)

	fmt.Println(http.StatusOK)

	checkResponseCode(t, http.StatusOK, testResponse.Code)
}
*/