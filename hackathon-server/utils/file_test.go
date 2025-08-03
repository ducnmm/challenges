package utils

import (
	"net/http"
	"testing"
)

func TestIsImageContentType(t *testing.T) {
	tests := []struct {
		contentType string
		expected    bool
	}{
		{"image/jpeg", true},
		{"image/jpg", true},
		{"image/png", true},
		{"image/gif", true},
		{"image/webp", true},
		{"image/bmp", true},
		{"image/svg+xml", true},
		{"text/plain", false},
		{"application/json", false},
		{"video/mp4", false},
		{"application/pdf", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsImageContentType(test.contentType)
		if result != test.expected {
			t.Errorf("IsImageContentType(%s) = %v, expected %v", 
				test.contentType, result, test.expected)
		}
	}
}

func TestGetClientIP(t *testing.T) {
	// Test with X-Forwarded-For header
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "192.168.1.1, 10.0.0.1")
	req.RemoteAddr = "127.0.0.1:8080"

	ip := GetClientIP(req)
	if ip != "192.168.1.1" {
		t.Errorf("Expected IP 192.168.1.1, got %s", ip)
	}

	// Test with X-Real-IP header
	req2, _ := http.NewRequest("GET", "/", nil)
	req2.Header.Set("X-Real-IP", "203.0.113.1")
	req2.RemoteAddr = "127.0.0.1:8080"

	ip2 := GetClientIP(req2)
	if ip2 != "203.0.113.1" {
		t.Errorf("Expected IP 203.0.113.1, got %s", ip2)
	}

	// Test with RemoteAddr only
	req3, _ := http.NewRequest("GET", "/", nil)
	req3.RemoteAddr = "198.51.100.1:9090"

	ip3 := GetClientIP(req3)
	if ip3 != "198.51.100.1" {
		t.Errorf("Expected IP 198.51.100.1, got %s", ip3)
	}
}