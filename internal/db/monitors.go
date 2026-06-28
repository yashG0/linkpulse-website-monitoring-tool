package db

import (
	"database/sql"

	"github.com/yashg0/linkpulse-website-monitoring-tool/internal/models"
)

func CreateMonitor(m *models.Monitor) error {
	result, err := DB.Exec(
		"INSERT INTO monitors(name, url, interval, enabled) VALUES(?, ?, ?, ?)",
		m.Name,
		m.URL,
		m.Interval,
		m.Enabled,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = int(id)
	return nil
}

func GetMonitor(id int) (*models.Monitor, error) {
	var m models.Monitor

	row := DB.QueryRow("SELECT id, name, url, interval, enabled FROM monitors WHERE id = ?", id)

	err := row.Scan(
		&m.ID,
		&m.Name,
		&m.URL,
		&m.Interval,
		&m.Enabled,
	)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func GetAllMonitors() ([]models.Monitor, error) {
	var monitors []models.Monitor
	rows, err := DB.Query("SELECT id, name, url, interval, enabled FROM monitors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var m models.Monitor
		err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.URL,
			&m.Interval,
			&m.Enabled,
		)
		if err != nil {
			return nil, err
		}

		monitors = append(monitors, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return monitors, nil
}

func UpdateMonitor(id int, m models.Monitor) error {
	result, err := DB.Exec("UPDATE monitors set name=?,url=?,interval=?,enabled=? WHERE id=?",
		m.Name, m.URL, m.Interval, m.Enabled, id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func DeleteMonitor(id int) error {
	result, err := DB.Exec("DELETE FROM monitors WHERE id = ?", id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return sql.ErrNoRows
	}
	return nil
}
