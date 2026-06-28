package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/config"
)

func main() {

	cfg := config.Load()

	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Working Done!")
	})

	log.Println("Server is running at", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
