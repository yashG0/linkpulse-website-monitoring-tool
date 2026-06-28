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
}
