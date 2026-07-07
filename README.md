# 🚀 LinkPulse – Website Monitoring Tool

LinkPulse is a backend website monitoring service built with **Go** and **SQLite** that continuously monitors website availability, records response metrics, and exposes a RESTful API for managing website monitors.

The project demonstrates Go concurrency using **goroutines**, background workers, and a clean modular architecture built primarily with Go's standard library.

---

## ✨ Features

- Website monitor CRUD operations
- Automatic background monitoring
- Configurable monitoring interval for each website
- Enable or disable monitors
- Track HTTP response time
- Record HTTP status codes
- Store monitoring history
- Concurrent monitoring using goroutines
- SQLite data persistence
- RESTful API
- Graceful server shutdown
- Clean and modular project architecture

---

## 🛠 Tech Stack

| Technology | Purpose |
|------------|---------|
| Go | Backend development |
| SQLite | Database |
| modernc.org/sqlite | SQLite Driver |
| net/http | HTTP Server |
| Goroutines | Background monitoring |
| RESTful API | API Architecture |

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

### Clone Repository

```bash
git clone https://github.com/yashG0/linkpulse.git
cd linkpulse
```

### Install Dependencies

```bash
go mod tidy
```

### Configure Environment

Create a `.env` file:

```env
PORT=:8080
DATABASE=linkpulse.db
```

### Run the Application

```bash
go run ./cmd/api
```

Server starts on:

```
http://localhost:8080
```

---

## 📡 API Endpoints

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | `/health` | Health Check |
| POST | `/api/monitors` | Create Monitor |
| GET | `/api/monitors` | Get All Monitors |
| GET | `/api/monitors/{id}` | Get Monitor |
| PUT | `/api/monitors/{id}` | Update Monitor |
| DELETE | `/api/monitors/{id}` | Delete Monitor |
| POST | `/api/monitors/{id}/check` | Trigger Manual Check |

### Example Request

```http
POST /api/monitors
```

```json
{
  "name": "Google",
  "url": "https://google.com",
  "interval": 60,
  "enabled": true
}
```

---

## 🗄 Database Schema

### monitors

| Column | Type |
|--------|------|
| id | INTEGER |
| name | TEXT |
| url | TEXT |
| interval | INTEGER |
| enabled | BOOLEAN |

### check_results

| Column | Type |
|--------|------|
| id | INTEGER |
| monitor_id | INTEGER |
| status_code | INTEGER |
| response_time | INTEGER |
| success | BOOLEAN |
| checked_at | DATETIME |
| error_message | TEXT |

---

## ⚙️ How It Works

```text
Create Monitor
        │
        ▼
Background Worker
        │
        ▼
Periodic Website Check
        │
        ▼
Measure Response Time
        │
        ▼
Store Status & History
        │
        ▼
Retrieve Monitoring Data via REST API
```

---

## 📸 Screenshots

### API Responses

> Add Postman screenshots

### Terminal Logs

> Add monitoring worker screenshots

### Database

> Add SQLite database screenshot

---

## 🎯 Why I Built This

I built **LinkPulse** to strengthen my understanding of backend development in Go, REST API design, concurrency using goroutines, and database-driven applications.

Instead of relying on third-party web frameworks, the project focuses on building maintainable backend services using Go's standard library and clean project organization.

---

## 🚀 Future Improvements

- Docker support
- JWT Authentication
- Worker pool for concurrent checks
- Email notifications
- Dashboard frontend
- Prometheus metrics
- PostgreSQL support
- Unit & Integration Tests
- Rate limiting
- Health analytics dashboard

---

## 👨‍💻 Author

**Yash Gaurkar**

- GitHub: https://github.com/yashG0
- LinkedIn: https://linkedin.com/in/yash-gaurkar-a897b3228
- Portfolio: https://yash-portfolio-app.vercel.app

---

⭐ If you found this project helpful, consider giving it a **Star**.
