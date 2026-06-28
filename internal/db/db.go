package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/config"
)

var DB *sql.DB

func Init(cfg config.Config) {
	db, err := sql.Open("sqlite", cfg.Database)

	if err != nil {
		log.Fatal("Failed to open database")
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB = db

	_, err = DB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatalf("Failed to enable foreign keys: %v", err)
	}

	createTables()
}

func createTables() {
	query := `
		CREATE TABLE IF NOT EXISTS monitors (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			url TEXT NOT NULL UNIQUE,
			interval INTEGER NOT NULL,
			enabled BOOLEAN NOT NULL
		);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create monitors table: %v", err)
	}

	query = `
		CREATE TABLE IF NOT EXISTS check_results (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			monitor_id INTEGER NOT NULL,
			status_code INTEGER NOT NULL,
			response_time INTEGER NOT NULL,
			success BOOLEAN NOT NULL,
			checked_at DATETIME NOT NULL,
			error_message TEXT,

			FOREIGN KEY (monitor_id) REFERENCES monitors(id) ON DELETE CASCADE
		);`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create check_results table: %v", err)
	}
}
