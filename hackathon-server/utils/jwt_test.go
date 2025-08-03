package utils

import (
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	secret := []byte("test-secret")
	userID := 123
	username := "testuser"

	// Test token generation
	token, err := GenerateToken(userID, username, secret)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Fatal("Token should not be empty")
	}

	// Test token validation
	claims, err := ValidateToken(token, secret)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected UserID %d, got %d", userID, claims.UserID)
	}

	if claims.Username != username {
		t.Errorf("Expected Username %s, got %s", username, claims.Username)
	}

	// Test token with wrong secret
	wrongSecret := []byte("wrong-secret")
	_, err = ValidateToken(token, wrongSecret)
	if err == nil {
		t.Error("Expected error when validating with wrong secret")
	}

	// Test invalid token format
	_, err = ValidateToken("invalid-token", secret)
	if err == nil {
		t.Error("Expected error when validating invalid token")
	}
}