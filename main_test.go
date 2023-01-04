package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServeApp(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.FileServer(http.Dir("./static"))
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("/ returned wrong status code, expected %v, got %v", http.StatusOK, rr.Code)
	}
	expectedString := "Welcome to the RHTE EMEA"
	if !strings.Contains(rr.Body.String(), expectedString) {
		t.Errorf("/ returned wrong body, expected %v, got %v", expectedString, rr.Body.String())
	}

}

func TestHealthEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthEndpoint)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("/ returned wrong status code, expected %v, got %v", http.StatusOK, rr.Code)
	}
	expectedString := "Healthy"
	if !strings.Contains(rr.Body.String(), expectedString) {
		t.Errorf("/ returned wrong body, expected %v, got %v", expectedString, rr.Body.String())
	}

}
