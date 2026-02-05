# Fiber REST API

Go + Fiber + GORM kullanılarak geliştirilmiş basit Todo API.

## Tech Stack
- Go
- Fiber v3
- GORM
- SQLite

## Endpoints

POST /todo
GET /todos
PUT /todos/:todoID
GET /todos/:todoID
DELETE /todos/:todoID

## Run
```bash
go run cmd/server/main.go
