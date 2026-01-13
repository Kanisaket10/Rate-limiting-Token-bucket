## Rate Limiting using Token Bucket (Go)

This project implements an API rate-limiting middleware in Go using the
Token Bucket algorithm.

### Features
- Token bucket rate limiting
- Thread-safe implementation using mutex
- Time-based token refill
- Returns HTTP 429 when limit exceeded

### How it Works
- Bucket starts with fixed capacity
- Each request consumes one token
- Tokens refill over time
- Requests are rejected when tokens are exhausted

### Run Locally
```bash
go run .
http://localhost:8080/ping

---