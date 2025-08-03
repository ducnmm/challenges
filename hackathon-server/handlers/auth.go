package handlers

import (
	"encoding/json"
	"hackathon-server/database"
	"hackathon-server/models"
	"hackathon-server/utils"
	"net/http"
	"strings"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendJSONResponse(w, http.StatusMethodNotAllowed, models.Response{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid JSON format",
		})
		return
	}

	if req.Username == "" || req.Password == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Username and password are required",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Error hashing password",
		})
		return
	}

	result, err := database.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)",
		req.Username, hashedPassword)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			utils.SendJSONResponse(w, http.StatusConflict, models.Response{
				Success: false,
				Message: "Username already exists",
			})
			return
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Error creating user",
		})
		return
	}

	userID, _ := result.LastInsertId()
	user := models.User{
		ID:       int(userID),
		Username: req.Username,
	}

	utils.SendJSONResponse(w, http.StatusCreated, models.Response{
		Success: true,
		Message: "User registered successfully",
		Data:    user,
	})
}

func LoginHandler(jwtSecret []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.SendJSONResponse(w, http.StatusMethodNotAllowed, models.Response{
				Success: false,
				Message: "Method not allowed",
			})
			return
		}

		var req models.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.SendJSONResponse(w, http.StatusBadRequest, models.Response{
				Success: false,
				Message: "Invalid JSON format",
			})
			return
		}

		var user models.User
		var hashedPassword string
		err := database.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?",
			req.Username).Scan(&user.ID, &user.Username, &hashedPassword)

		if err != nil {
			utils.SendJSONResponse(w, http.StatusUnauthorized, models.Response{
				Success: false,
				Message: "Invalid credentials",
			})
			return
		}

		if !utils.CheckPasswordHash(req.Password, hashedPassword) {
			utils.SendJSONResponse(w, http.StatusUnauthorized, models.Response{
				Success: false,
				Message: "Invalid credentials",
			})
			return
		}

		token, err := utils.GenerateToken(user.ID, user.Username, jwtSecret)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusInternalServerError, models.Response{
				Success: false,
				Message: "Error generating token",
			})
			return
		}

		utils.SendJSONResponse(w, http.StatusOK, models.Response{
			Success: true,
			Message: "Login successful",
			Data: models.LoginResponse{
				Token: token,
				User:  user,
			},
		})
	}
}