package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/checker"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/db"
)

func CheckerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	// Validate the path before accessing parts[2]
	if len(parts) != 4 || parts[3] != "check" {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid monitor ID", http.StatusBadRequest)
		return
	}

	m, err := db.GetMonitor(id)
	if err == sql.ErrNoRows {
		http.Error(w, "Monitor not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	result := checker.CheckMonitor(m)

	if err := db.CreateResult(&result); err != nil {
		http.Error(w, "Failed to save check result", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
