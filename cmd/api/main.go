package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/checker"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/config"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/db"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/server"
)

func main() {
	// capture the Ctrl+C signal from user and stop it when function end!
	ctx, stopSignals := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopSignals()

	// either the project configuration load or gracefully exit the program!
	cfg := config.MustLoad()

	// init the sqlite db instance and exit when the function end!
	db.Init(cfg)
	defer db.DB.Close()
	// create new server router/mux with configuration!
	srv := server.New(cfg)

	// added waitgroup to track the no of goroutine by adding and subtracting as per the work done!
	var wg sync.WaitGroup

	// start the workers here
	wg.Go(func() {
		checker.StartWorker(ctx.Done())
	})

	// listen the server over go-routine function and exit the program when error!
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// wait, untill the signal arrives
	<-ctx.Done()

	// handle the second ctrl+c force quit
	stopSignals()
	
	// gracefully shutdown the server and docker
	log.Println("Shutdown signal received")
	shutDownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// handlling the shutdown here
	err := srv.Shutdown(shutDownCtx)
	if err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	// wait untill all goroutine finished task.
	wg.Wait()
	log.Println("Application stopped")
}
