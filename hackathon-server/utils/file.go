package utils

import (
	"net/http"
	"strings"
)

func IsImageContentType(contentType string) bool {
	imageTypes := []string{
		"image/jpeg",
		"image/jpg", 
		"image/png",
		"image/gif",
		"image/webp",
		"image/bmp",
		"image/svg+xml",
	}

	for _, imageType := range imageTypes {
		if contentType == imageType {
			return true
		}
	}
	return false
}

func GetClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		return strings.Split(xForwardedFor, ",")[0]
	}

	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	return strings.Split(r.RemoteAddr, ":")[0]
}