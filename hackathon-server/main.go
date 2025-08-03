package main

import (
	"fmt"
	"hackathon-server/database"
	"hackathon-server/handlers"
	"hackathon-server/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	// Configuration from environment variables
	jwtSecret := []byte(getEnv("JWT_SECRET", "your-secret-key"))
	port := getEnv("PORT", "8080")
	dbPath := getEnv("DB_PATH", "./hackathon.db")

	// Initialize database
	if err := database.InitDB(dbPath); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.Close()

	// Setup routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler(jwtSecret))
	http.HandleFunc("/upload", middleware.AuthMiddleware(jwtSecret)(handlers.UploadHandler))

	// Start server
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Printf("Database: %s\n", dbPath)
	fmt.Println("Open http://localhost:" + port + " to test the application")
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}