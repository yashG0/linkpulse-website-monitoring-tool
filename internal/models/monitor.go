package models

type Monitor struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Interval int    `json:"interval"`
	Enabled  bool   `json:"enabled"`
}
