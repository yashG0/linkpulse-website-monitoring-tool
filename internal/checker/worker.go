package checker

import (
	"database/sql"
	"log"
	"time"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/db"
)

func StartWorker() {
	for {
		monitors, err := db.GetEnabledMonitors()
		if err != nil {
			log.Printf("Failed to fetch monitors: %v", err)
			continue
		}

		for _, m := range monitors {
			lastChecked, err := db.GetLastCheckTime(m.ID)
			if err == sql.ErrNoRows {
				result := CheckMonitor(&m)
				if err := db.CreateResult(&result); err != nil {
					log.Println(err)
				}
				continue
			}
			if err != nil {
				log.Println(err)
				continue
			}

			elapsed := time.Since(lastChecked)
			if elapsed < time.Duration(m.Interval)*time.Second {
				continue
			}

			result := CheckMonitor(&m)
			err = db.CreateResult(&result)
			if err != nil {
				log.Print(err)
				continue
			}
		}
		time.Sleep(5 * time.Second)
	}
}
