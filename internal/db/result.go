package db

import (
	"time"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/models"
)

func CreateResult(r *models.CheckResult) error {
	result, err := DB.Exec(
		"INSERT INTO check_results(monitor_id, status_code, response_time, success, checked_at, error_message) VALUES(?, ?, ?, ?, ?, ?)",
		r.MonitorID,
		r.StatusCode,
		r.ResponseTime,
		r.Success,
		r.CheckedAt,
		r.ErrorMessage,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	r.ID = int(id)
	return nil
}

func GetResult(id int) (*models.CheckResult, error) {
	var r models.CheckResult

	row := DB.QueryRow("SELECT id, monitor_id, status_code, response_time, success, checked_at, error_message FROM check_results WHERE id = ?", id)

	err := row.Scan(
		&r.ID,
		&r.MonitorID,
		&r.StatusCode,
		&r.ResponseTime,
		&r.Success,
		&r.CheckedAt,
		&r.ErrorMessage,
	)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func GetAllResults() ([]models.CheckResult, error) {
	var results []models.CheckResult
	rows, err := DB.Query("SELECT id, monitor_id, status_code, response_time, success, checked_at, error_message FROM check_results")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var r models.CheckResult
		err := rows.Scan(
			&r.ID,
			&r.MonitorID,
			&r.StatusCode,
			&r.ResponseTime,
			&r.Success,
			&r.CheckedAt,
			&r.ErrorMessage,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func GetLastCheckTime(monitorID int) (time.Time, error) {
	var lastChecked time.Time
	row := DB.QueryRow("SELECT checked_at FROM check_results WHERE monitor_id=? ORDER BY checked_at DESC LIMIT 1", monitorID)

	err := row.Scan(&lastChecked)
	if err != nil {
		return time.Time{}, err
	}
	return lastChecked, nil
}
