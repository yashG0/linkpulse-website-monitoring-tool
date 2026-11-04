package checker

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/db"
)

func StartWorker(stop <-chan struct{}) {
	for {
		checkDue(stop)

		select {
		case <-time.After(5 * time.Second):
		case <-stop:
			log.Println("Worker stopped!")
			return
		}
	}
}

func checkDue(stop <-chan struct{}) {
	monitors, err := db.GetEnabledMonitors()
	if err != nil {
		log.Printf("Failed to fetch monitor: %v", err)
		return
	}

	for _, m := range monitors {
		select {
		case <-stop:
			return
		default:
		}

		var lastChecked, err = db.GetLastCheckTime(m.ID)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
			continue
		}
		if err == nil && time.Since(lastChecked) < time.Duration(m.Interval)*time.Second {
			continue
		}
		result := CheckMonitor(&m)
		err = db.CreateResult(&result)
		if err != nil {
			log.Println(err)
		}
	}
}
