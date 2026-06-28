package server

import (
	"log"
	"net/http"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/config"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/handlers"
)

func Start(cfg config.Config) {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/api/monitors", handlers.MonitorsHandler)
	mux.HandleFunc("/api/monitors/", handlers.MonitorHandler)
	
	server := &http.Server{
		Addr:    cfg.Port,
		Handler: mux,
	}

	log.Println("Server is running at", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
