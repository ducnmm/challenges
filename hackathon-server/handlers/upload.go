package handlers

import (
	"fmt"
	"hackathon-server/database"
	"hackathon-server/models"
	"hackathon-server/utils"
	"io"
	"net/http"
	"os"
	"time"
)

const MaxUploadSize = int64(8 << 20) // 8MB

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendJSONResponse(w, http.StatusMethodNotAllowed, models.Response{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	err := r.ParseMultipartForm(MaxUploadSize)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, models.Response{
			Success: false,
			Message: "File too large or invalid form data",
		})
		return
	}

	file, fileHeader, err := r.FormFile("data")
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, models.Response{
			Success: false,
			Message: "No file found in 'data' field",
		})
		return
	}
	defer file.Close()

	if fileHeader.Size > MaxUploadSize {
		utils.SendJSONResponse(w, http.StatusBadRequest, models.Response{
			Success: false,
			Message: "File size exceeds 8MB limit",
		})
		return
	}

	contentType := fileHeader.Header.Get("Content-Type")
	if !utils.IsImageContentType(contentType) {
		utils.SendJSONResponse(w, http.StatusBadRequest, models.Response{
			Success: false,
			Message: "File must be an image",
		})
		return
	}

	userIDStr := r.Header.Get("X-User-ID")
	if userIDStr == "" {
		utils.SendJSONResponse(w, http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "User ID not found",
		})
		return
	}

	userID := 0
	fmt.Sscanf(userIDStr, "%d", &userID)

	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d_%s", timestamp, fileHeader.Filename)
	filepath := fmt.Sprintf("/tmp/%s", filename)

	dst, err := os.Create(filepath)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Error creating file",
		})
		return
	}
	defer dst.Close()

	fileSize, err := io.Copy(dst, file)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Error saving file",
		})
		return
	}

	metadata := models.FileMetadata{
		Filename:    fileHeader.Filename,
		ContentType: contentType,
		Size:        fileSize,
		Path:        filepath,
		UserID:      userID,
		UploadedAt:  time.Now(),
		RemoteAddr:  utils.GetClientIP(r),
		UserAgent:   r.UserAgent(),
	}

	result, err := database.DB.Exec(`
		INSERT INTO file_metadata 
		(filename, content_type, size, path, user_id, remote_addr, user_agent) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		metadata.Filename, metadata.ContentType, metadata.Size,
		metadata.Path, metadata.UserID, metadata.RemoteAddr, metadata.UserAgent)

	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Error saving file metadata",
		})
		return
	}

	metadataID, _ := result.LastInsertId()
	metadata.ID = int(metadataID)

	utils.SendJSONResponse(w, http.StatusOK, models.Response{
		Success: true,
		Message: "File uploaded successfully",
		Data:    metadata,
	})
}