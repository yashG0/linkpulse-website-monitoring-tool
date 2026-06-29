package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/db"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/models"
)

func getMonitor(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 {
		http.NotFound(w, r)
		return
	}

	idInt, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid monitor ID", http.StatusBadRequest)
		return
	}

	monitor, err := db.GetMonitor(idInt)
	if err == sql.ErrNoRows {
		http.Error(w, "Monitor not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Failed to fetched monitor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(monitor)
}

func createMonitor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var m models.Monitor
	deco := json.NewDecoder(r.Body)
	deco.DisallowUnknownFields()
	err := deco.Decode(&m)

	if err != nil {
		http.Error(w, "Failed to read value", http.StatusBadRequest)
		return
	}
	if m.Name == "" || m.URL == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	err = db.CreateMonitor(&m)
	if err != nil {
		http.Error(w, "Failed to create monitor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

func getAllMonitors(w http.ResponseWriter, r *http.Request) {
	monitors, err := db.GetAllMonitors()
	if err != nil {
		http.Error(w, "Failed to fetched data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(monitors)
}

func updateMonitor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 {
		http.NotFound(w, r)
		return
	}

	idInt, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid monitor ID", http.StatusBadRequest)
		return
	}

	var newMonitor models.Monitor

	deco := json.NewDecoder(r.Body)
	deco.DisallowUnknownFields()
	err = deco.Decode(&newMonitor)

	if err != nil {
		http.Error(w, "Failed to read value", http.StatusBadRequest)
		return
	}

	if newMonitor.Name == "" || newMonitor.URL == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	err = db.UpdateMonitor(idInt, newMonitor)
	if err == sql.ErrNoRows {
		http.Error(w, "Monitor not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteMonitor(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 {
		http.NotFound(w, r)
		return
	}

	idInt, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid monitor ID", http.StatusBadRequest)
		return
	}

	err = db.DeleteMonitor(idInt)
	if err == sql.ErrNoRows {
		http.Error(w, "Monitor not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func MonitorHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/check") {
		CheckerHandler(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getMonitor(w, r)
	case http.MethodPut:
		updateMonitor(w, r)
	case http.MethodDelete:
		deleteMonitor(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
