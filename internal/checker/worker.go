package checker

import (
	"log"
	"time"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/db"
)

func StartWorker() {
	for {
		monitors, err := db.GetEnabledMonitors()
		if err != nil {
			continue
		}

		for _, m := range monitors {
			result := CheckMonitor(&m)
			err := db.CreateResult(&result)
			if err != nil {
				log.Print(err)
				continue
			}
		}
		time.Sleep(1 * time.Minute)
	}
}
