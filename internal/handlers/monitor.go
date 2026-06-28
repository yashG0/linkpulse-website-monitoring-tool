package handlers

import "net/http"

func getMonitor(w http.ResponseWriter, r *http.Request) {
	
}

func createMonitor(w http.ResponseWriter, r *http.Request) {
	
}

func getAllMonitors(w http.ResponseWriter, r *http.Request) {

}

func updateMonitor(w http.ResponseWriter, r *http.Request) {

}
func deleteMonitor(w http.ResponseWriter, r *http.Request) {

}

func MonitorHandler(w http.ResponseWriter, r *http.Request) {
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

func MonitorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAllMonitors(w, r)
	case http.MethodPost:
		createMonitor(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
