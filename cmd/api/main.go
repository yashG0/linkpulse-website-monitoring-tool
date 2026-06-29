package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/checker"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/config"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/db"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/server"
)

func main() {
	stop := make(chan struct{})
	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer signal.Stop(quit)
	cfg := config.Load()

	db.Init(cfg)
	defer db.DB.Close()

	srv := server.New(cfg)
	go checker.StartWorker(stop)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-quit
	log.Println("Shutdown signal received")
	close(stop)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	log.Println("Application stopped")
}
