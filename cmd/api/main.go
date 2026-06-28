package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Working Done!")
	})
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not set")
	}

	log.Println("Server is running at", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
