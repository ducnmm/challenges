package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	if err := createTables(); err != nil {
		return fmt.Errorf("error creating tables: %v", err)
	}

	fmt.Println("Database initialized successfully")
	return nil
}

func createTables() error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	createFilesTable := `
	CREATE TABLE IF NOT EXISTS file_metadata (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		filename TEXT NOT NULL,
		content_type TEXT NOT NULL,
		size INTEGER NOT NULL,
		path TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		remote_addr TEXT,
		user_agent TEXT,
		FOREIGN KEY (user_id) REFERENCES users (id)
	);`

	if _, err := DB.Exec(createUsersTable); err != nil {
		return err
	}

	if _, err := DB.Exec(createFilesTable); err != nil {
		return err
	}

	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}