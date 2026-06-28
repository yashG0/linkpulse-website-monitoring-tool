package models

import "time"

type CheckResult struct {
	ID           int       `json:"id"`
	MonitorID    int       `json:"monitor_id"`
	StatusCode   int       `json:"status_code"`
	ResponseTime int       `json:"response_time"`
	Success      bool      `json:"success"`
	CheckedAt    time.Time `json:"checked_at"`
	ErrorMessage string    `json:"error_message"`
}
