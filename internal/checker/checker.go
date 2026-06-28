package checker

import (
	"net/http"
	"time"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/models"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func CheckMonitor(m *models.Monitor) models.CheckResult {
	var r models.CheckResult
	r.MonitorID = m.ID
	r.CheckedAt = time.Now()

	start := time.Now()

	resp, err := client.Get(m.URL)

	if err != nil {
		r.ErrorMessage = err.Error()
		r.Success = false
		return r
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	r.StatusCode = resp.StatusCode
	r.ResponseTime = int(elapsed.Milliseconds())
	r.Success = resp.StatusCode >= 200 && resp.StatusCode < 400

	return r
}
