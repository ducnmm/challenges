package utils

import "testing"

func TestHashAndCheckPassword(t *testing.T) {
	password := "testpassword123"

	// Test password hashing
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if hash == "" {
		t.Fatal("Hash should not be empty")
	}

	if hash == password {
		t.Error("Hash should be different from original password")
	}

	// Test password checking - correct password
	if !CheckPasswordHash(password, hash) {
		t.Error("Password should match hash")
	}

	// Test password checking - wrong password
	wrongPassword := "wrongpassword"
	if CheckPasswordHash(wrongPassword, hash) {
		t.Error("Wrong password should not match hash")
	}

	// Test empty password
	emptyHash, err := HashPassword("")
	if err != nil {
		t.Fatalf("Failed to hash empty password: %v", err)
	}

	if !CheckPasswordHash("", emptyHash) {
		t.Error("Empty password should match its hash")
	}
}