# Weather API - Go Fiber Application

A high-performance weather query application built with Go and Fiber framework that aggregates data from multiple weather services.

## Features

- **Request Batching**: Groups multiple requests for the same location within a 5-second window
- **Multi-Service Aggregation**: Fetches data from two weather services and returns the average temperature
- **Non-Blocking Database Logging**: Asynchronous database writes don't delay user responses
- **Auto-Execution**: Batches execute immediately when 10 requests are queued or after 5-second timeout
- **MariaDB/MySQL Support**: Persistent storage of all weather queries
- **Concurrent API Calls**: Parallel requests to both weather services for faster response

## Architecture

```
weather-api/
├── cmd/
│   └── server/
│       └── main.go             
├── internal/  
│   ├── middleware/
│   │   └── logger.go       
│   ├── dto/
│   │   └── weather_response.go       
│   ├── route/
│   │   ├── routes.go           
│   │   ├── health.go           
│   │   └── weather.go      
│   ├── handler/
│   │   ├── health.go           
│   │   └── weather.go           
│   ├── orchestrator/
│   │   └── weather.go           
│   ├── service/
│   │   ├── weather.go           
│   │   └── batch.go             
│   ├── model/
│   │   └── batch.go           
│   ├── weather/
│   │   └── weather.go           
│   ├── database/
│   │   └── connection.go         
│   ├── repository/
│   │   └── weather_query.go         
│   └── pkg/
│       ├── weatherapi/    
│       │   ├── dto.go             
│       │   └── client.go             
│       └── weatherstack/    
│           ├── dto.go             
│           └── client.go             
├── configs/
├── .env.example                 
├── go.mod
└── README.md
```

## Prerequisites

- Go 1.24 or higher
- MariaDB 10.5+ or MySQL 8.0+
- Weather API keys (included in `.env.example`)

## Installation

1. **Clone or extract the project**:
   ```bash
   cd weather-api
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Set up MariaDB database**:
   ```sql
   CREATE DATABASE weather-db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
   CREATE USER 'weather_user'@'localhost' IDENTIFIED BY 'your_password';
   GRANT ALL PRIVILEGES ON weather-db.* TO 'weather_user'@'localhost';
   FLUSH PRIVILEGES;
   ```

4. **Configure environment variables**:
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

5. **Run the application**:
   ```bash
   go run cmd/server/main.go
   ```

   Or build and run:
   ```bash
   go build -o bin/weather-api cmd/server/main.go
   ./bin/weather-api
   ```

## Configuration

Edit the `.env` file:

```env
# Server
PORT=8080
ENV=development

# Database
DB_HOST=localhost
DB_PORT=3307
DB_USER=weather_user
DB_PASSWORD=your_password
DB_NAME=weather-db

# Weather APIs
WEATHERAPI_KEY=123456..
WEATHERSTACK_KEY=123456...

# Batching
MAX_BATCH_SIZE=10
BATCH_TIMEOUT_SECONDS=5
```

## API Endpoints

### Get Weather
```http
GET /weather?q=<location>
```

**Parameters**:
- `q` (required): Location name (e.g., "Istanbul", "London", "New York")

**Response** (200 OK):
```json
{
  "location": "Istanbul",
  "temperature": 15.5
}
```

**Error Response** (400 Bad Request):
```json
{
  "error": "location query parameter 'q' is required"
}
```

### Health Check
```http
GET /health
```

**Response**:
```json
{
  "status": "ok",
  "message": "Weather API is running"
}
```

## Usage Examples

### Single Request
```bash
curl "http://localhost:8080/weather?q=Istanbul"
```

Response time: ~6 seconds (5s wait + 1s API call)

### Multiple Concurrent Requests
```bash
curl "http://localhost:8080/weather?q=Istanbul"

curl "http://localhost:8080/weather?q=Istanbul"

curl "http://localhost:8080/weather?q=Istanbul"
```

All three requests will be batched and receive the same response after a single API call.

### Testing Batch Limits
```bash
for i in {1..10}; do
  curl "http://localhost:8080/weather?q=Istanbul" &
done
wait
```

## Request Batching Logic

1. **First Request**: Starts a 5-second timer
2. **Subsequent Requests**: Join the existing batch
3. **Execution Triggers**:
   - 5-second timeout expires
   - 10 requests queued (max batch size)
4. **Response**: All batched requests receive the same result
5. **Database Logging**: Happens asynchronously after response sent

## Database Schema

```sql
CREATE TABLE weather_queries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    location VARCHAR(255) NOT NULL,
    service_1_temperature FLOAT NOT NULL,
    service_2_temperature FLOAT NOT NULL,
    request_count INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_location (location),
    INDEX idx_created_at (created_at)
);
```

## Running Tests

```bash
go test ./...

go test -cover ./...

go test -v ./...

go test ./internal/services/
go test ./internal/handlers/
```

## Docker Support (Optional)

Create a `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o weather-api cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/weather-api .
COPY .env .
EXPOSE 8080
CMD ["./weather-api"]
```

Build and run:
```bash
docker build -t weather-api .
docker run -p 8080:8080 --env-file .env weather-api
```

## Performance Considerations

- **Concurrent API Calls**: Both weather services are queried in parallel
- **Connection Pooling**: Database connection pool configured (25 max, 5 idle)
- **Async Logging**: Database writes don't block HTTP responses
- **Batch Processing**: Reduces API calls by up to 10x during high traffic


## Development

### Project Structure Best Practices
- `cmd/`: Application entry points
- `internal/`: Private application code
- `pkg/`: Public libraries (if needed)
- `configs/`: Configuration files

### Adding New Features
1. Add models in `internal/models/`
2. Implement business logic in `internal/services/`
3. Create handlers in `internal/handlers/`
4. Add routes in `cmd/server/main.go`
5. Write tests for all new code
