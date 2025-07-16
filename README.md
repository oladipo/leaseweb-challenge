# Leaseweb Challenge API

A RESTful API built with Go and Gin for managing a server inventory, including filtering and paginated listing. The app demonstrates best practices in Go web development, middleware usage, and clean architecture.

---

## Hosted Application

You can access the live application here: [http://ec2-54-88-56-9.compute-1.amazonaws.com](http://ec2-54-88-56-9.compute-1.amazonaws.com)

---

## Features
- List all servers with pagination
- Filter servers by hardware and location
- Prometheus metrics
- Rate limiting, security, and compression middleware

---

## Getting Started

### Prerequisites
- Go 1.18+
- PostgreSQL (or compatible database)

### Clone the repository
```bash
git clone https://github.com/oladipo/leaseweb-challenge.git
cd leaseweb-challenge
```

### Configuration
Edit the config file or set environment variables as needed. Example config is in `internal/config/config.go`.

### Database Setup
1. Create a PostgreSQL database (e.g., `leaseweb`):
   ```bash
   createdb leaseweb
   ```
2. Apply the SQL schema (see `sql/init.sql` if present):
   ```bash
   psql leaseweb < sql/init.sql
   ```
3. Update your DB credentials in the config/environment variables.

### Install dependencies
```bash
go mod tidy
```

### Run with Docker Compose
You can run the entire stack (API, PostgreSQL, Redis) using Docker Compose:

```bash
docker-compose up --build
```

This will start the API on `localhost:8080`, PostgreSQL on `localhost:5432`, and Redis on `localhost:6379`.

You can stop the services with:
```bash
docker-compose down
```

---

### Run the application (locally)
If you prefer to run the Go app locally (without Docker):
```bash
go run cmd/server/main.go
```

The server will start on `localhost:8080` by default.

---

## Frontend Application

A modern React app is provided in the `frontend` directory for searching and filtering servers via the API.

### Features
- Web form for filtering by Storage (range slider), RAM (checkboxes), Harddisk type (dropdown), and Location (dropdown)
- Results displayed in a responsive table
- Fast development with Vite
- Dockerized for unified startup with backend

### Local Development
```bash
cd frontend
npm install
npm run dev
```
- App runs at [http://localhost:5173](http://localhost:5173) and proxies API requests to the backend.

### Build & Run with Docker
```bash
docker build -t leaseweb-frontend ./frontend
docker run -p 3000:80 leaseweb-frontend
```
- App available at [http://localhost:3000](http://localhost:3000)

### Unified Startup (Docker Compose)
From the project root:
```bash
docker-compose up --build
```
- Backend: [http://localhost:8080](http://localhost:8080)
- Frontend: [http://localhost:3000](http://localhost:3000)

---

## API Documentation

### List Servers
- **Endpoint:** `GET /servers`
- **Query parameters:**
  - `page` (optional, integer): Page number (default: 0)
  - `limit` (optional, integer): Items per page (default: 0 = all)
- **Response:**
```json
{
  "data": [
    {
      "id": "1",
      "model": "Dell R210Intel Xeon X3440",
      "ram": "16GB",
      "hdd": "2x2TB SATA2",
      "location": "AmsterdamAMS-01",
      "price": "$49.99"
    },
    ...
  ]
}
```

### Filter Servers
- **Endpoint:** `POST /servers/filter`
- **Body:**
```json
{
  "storage": "2TB",   // optional
  "ram": "16GB",      // optional
  "hdd": "SATA2",     // optional
  "location": "AMS-01"// optional
}
```
- **Response:**
```json
{
  "count": 1,
  "data": [
    {
      "id": "1",
      "model": "Dell R210Intel Xeon X3440",
      "ram": "16GB",
      "hdd": "2x2TB SATA2",
      "location": "AmsterdamAMS-01",
      "price": "$49.99"
    }
  ]
}
```

---

## Middleware
- **CORS**: Enabled for all origins
- **GZIP**: Compression enabled
- **RequestID**: Each request gets a unique ID
- **Security Headers**: Uses `unrolled/secure`
- **Rate Limiting**: 60 requests/minute (in-memory)
- **Prometheus Metrics**: `/metrics` endpoint

---

## Testing
Run unit tests with:
```bash
go test ./...
go test ./internal/handlers/... -v
```

---
