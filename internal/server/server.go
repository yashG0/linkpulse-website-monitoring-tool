package server

import (
	"net/http"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/config"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/handlers"
)

func New(cfg config.Config) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/api/monitors", handlers.MonitorsHandler)
	mux.HandleFunc("/api/monitors/", handlers.MonitorHandler)
	server := &http.Server{
		Addr:    cfg.Port,
		Handler: mux,
	}
	return server
}
