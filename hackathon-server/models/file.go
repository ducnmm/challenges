package models

import "time"

type FileMetadata struct {
	ID          int       `json:"id"`
	Filename    string    `json:"filename"`
	ContentType string    `json:"content_type"`
	Size        int64     `json:"size"`
	Path        string    `json:"path"`
	UserID      int       `json:"user_id"`
	UploadedAt  time.Time `json:"uploaded_at"`
	RemoteAddr  string    `json:"remote_addr"`
	UserAgent   string    `json:"user_agent"`
}