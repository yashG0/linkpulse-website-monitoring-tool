package main

import (
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/checker"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/config"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/db"
	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/server"
)

func main() {

	cfg := config.Load()

	db.Init(cfg)
	defer db.DB.Close()
	go checker.StartWorker()
	server.Start(cfg)
}
