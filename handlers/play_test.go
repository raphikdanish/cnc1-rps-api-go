package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Helper function to perform HTTP request
func performRequest(body map[string]string) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/play", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(PlayHandler)
	handler.ServeHTTP(rr, req)

	return rr
}

// Test: valid input
func TestPlayHandler_ValidInput(t *testing.T) {
	body := map[string]string{
		"choice": "rock",
	}

	rr := performRequest(body)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	var resp map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("could not parse response: %v", err)
	}

	if resp["player_choice"] != "rock" {
		t.Errorf("expected player_choice rock, got %s", resp["player_choice"])
	}

	if resp["computer_choice"] == "" {
		t.Errorf("expected computer_choice to be set")
	}

	if resp["result"] == "" {
		t.Errorf("expected result to be set")
	}
}

// Test: invalid choice
func TestPlayHandler_InvalidChoice(t *testing.T) {
	body := map[string]string{
		"choice": "banana",
	}

	rr := performRequest(body)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rr.Code)
	}
}

// Test: invalid JSON
func TestPlayHandler_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/play", bytes.NewBuffer([]byte("invalid-json")))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(PlayHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rr.Code)
	}
}
