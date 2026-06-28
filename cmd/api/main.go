package main

import (
	"log"
	"net/http"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/config"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/handlers"
)

func main() {

	cfg := config.Load()

	http.HandleFunc("/api/test", handlers.TestHandler)

	log.Println("Server is running at", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
