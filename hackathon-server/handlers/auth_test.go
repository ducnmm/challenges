package handlers

import (
	"bytes"
	"encoding/json"
	"hackathon-server/database"
	"hackathon-server/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Setup test database
	database.InitDB(":memory:")
	code := m.Run()
	database.Close()
	os.Exit(code)
}

func TestRegisterHandler(t *testing.T) {
	// Test successful registration
	reqBody := models.RegisterRequest{
		Username: "testuser",
		Password: "password123",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	RegisterHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var response models.Response
	json.NewDecoder(w.Body).Decode(&response)

	if !response.Success {
		t.Error("Expected success to be true")
	}

	// Test duplicate username
	req2 := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()

	RegisterHandler(w2, req2)

	if w2.Code != http.StatusConflict {
		t.Errorf("Expected status 409, got %d", w2.Code)
	}

	// Test invalid JSON
	req3 := httptest.NewRequest("POST", "/register", bytes.NewBuffer([]byte("invalid json")))
	req3.Header.Set("Content-Type", "application/json")
	w3 := httptest.NewRecorder()

	RegisterHandler(w3, req3)

	if w3.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w3.Code)
	}

	// Test missing username/password
	emptyReq := models.RegisterRequest{Username: "", Password: ""}
	emptyJson, _ := json.Marshal(emptyReq)
	req4 := httptest.NewRequest("POST", "/register", bytes.NewBuffer(emptyJson))
	req4.Header.Set("Content-Type", "application/json")
	w4 := httptest.NewRecorder()

	RegisterHandler(w4, req4)

	if w4.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w4.Code)
	}
}

func TestLoginHandler(t *testing.T) {
	secret := []byte("test-secret")

	// First register a user
	reqBody := models.RegisterRequest{
		Username: "logintest",
		Password: "password123",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	RegisterHandler(w, req)

	// Test successful login
	loginReq := models.LoginRequest{
		Username: "logintest",
		Password: "password123",
	}
	loginJson, _ := json.Marshal(loginReq)
	req2 := httptest.NewRequest("POST", "/login", bytes.NewBuffer(loginJson))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()

	LoginHandler(secret)(w2, req2)

	if w2.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w2.Code)
	}

	var response models.Response
	json.NewDecoder(w2.Body).Decode(&response)

	if !response.Success {
		t.Error("Expected success to be true")
	}

	// Test wrong password
	wrongLoginReq := models.LoginRequest{
		Username: "logintest",
		Password: "wrongpassword",
	}
	wrongJson, _ := json.Marshal(wrongLoginReq)
	req3 := httptest.NewRequest("POST", "/login", bytes.NewBuffer(wrongJson))
	req3.Header.Set("Content-Type", "application/json")
	w3 := httptest.NewRecorder()

	LoginHandler(secret)(w3, req3)

	if w3.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w3.Code)
	}

	// Test non-existent user
	nonExistentReq := models.LoginRequest{
		Username: "nonexistent",
		Password: "password123",
	}
	nonExistentJson, _ := json.Marshal(nonExistentReq)
	req4 := httptest.NewRequest("POST", "/login", bytes.NewBuffer(nonExistentJson))
	req4.Header.Set("Content-Type", "application/json")
	w4 := httptest.NewRecorder()

	LoginHandler(secret)(w4, req4)

	if w4.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w4.Code)
	}
}