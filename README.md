# 🚀 LinkPulse – Website Monitoring Tool

LinkPulse is a lightweight website monitoring service built with **Go** and **SQLite**. It periodically checks websites, records their status and response time, and exposes a clean REST API for managing monitors.

---

## ✨ Features

* Create, update and delete website monitors
* Enable or disable monitoring
* Automatic background worker
* Custom monitoring interval for each website
* Measure response time
* Store HTTP status codes
* Save monitoring history
* SQLite database
* RESTful API
* Graceful shutdown
* Clean project structure using Go's standard library

---

## 🛠️ Tech Stack

* Go
* SQLite
* modernc.org/sqlite
* net/http
* Goroutines
* REST API

---

## 📂 Project Structure

```text
cmd/
└── api/
    └── main.go

internal/
├── checker/
├── config/
├── db/
├── handlers/
├── models/
└── server/
```

---

## ⚙️ Getting Started

### Clone the repository

```bash
git clone https://github.com/yourusername/linkpulse-website-monitoring-tool.git
cd linkpulse-website-monitoring-tool
```

### Install dependencies

```bash
go mod tidy
```

### Create `.env`

```env
PORT=:8080
DATABASE=linkpulse.db
```

### Run

```bash
go run ./cmd/api
```

---

## 📡 API Endpoints

### Health Check

```
GET /health
```

---

### Create Monitor

```
POST /api/monitors
```

Example JSON

```json
{
  "name": "Google",
  "url": "https://google.com",
  "interval": 60,
  "enabled": true
}
```

---

### Get All Monitors

```
GET /api/monitors
```

---

### Get Monitor

```
GET /api/monitors/{id}
```

---

### Update Monitor

```
PUT /api/monitors/{id}
```

---

### Delete Monitor

```
DELETE /api/monitors/{id}
```

---

### Check Website Manually

```
POST /api/monitors/{id}/check
```

---

## 🗄 Database

### monitors

| Column   | Type    |
| -------- | ------- |
| id       | INTEGER |
| name     | TEXT    |
| url      | TEXT    |
| interval | INTEGER |
| enabled  | BOOLEAN |

### check_results

| Column        | Type     |
| ------------- | -------- |
| id            | INTEGER  |
| monitor_id    | INTEGER  |
| status_code   | INTEGER  |
| response_time | INTEGER  |
| success       | BOOLEAN  |
| checked_at    | DATETIME |
| error_message | TEXT     |

---

## 🔄 How It Works

1. User creates a monitor.
2. Background worker loads all enabled monitors.
3. Each monitor is checked based on its configured interval.
4. HTTP status code, response time and timestamp are stored in SQLite.
5. Monitoring history can be retrieved through the API.

---

## 📸 Screenshots

> Add screenshots of your API responses, terminal logs and project structure here.

---

## 🔮 Future Improvements

* Docker support
* Unit tests
* Worker pool for concurrent checks
* Email notifications
* Dashboard frontend
* Prometheus metrics
* Authentication
* PostgreSQL support

---

## 👨‍💻 Author

**Yash**

GitHub: https://github.com/yashG0

---

⭐ If you found this project useful, consider giving it a star!
