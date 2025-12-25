# Copilot Instructions for go-blog-api

## Project Overview

This is a learning-focused Go web API project implementing a three-tier architecture (API/Service/Repository) for a blog backend. The project uses Go standard library (`net/http`) with plans to add GORM for database operations.

## Architecture & Structure

### Three-Tier Pattern

- **API Layer** (`internal/api/`): HTTP handlers, request/response handling. Example: `health.go` uses service layer and returns JSON
- **Service Layer** (`internal/service/`): Business logic. Example: `HealthService` contains methods like `Get()` returning domain objects
- **Repository Layer** (`internal/repository/`): Data access (currently placeholder, future GORM integration)
- **Model Layer** (`internal/model/`): Domain/database models (GORM structs)

### Module Import Pattern

Use the module name `go-blog-api` as import prefix:

```go
import (
    "go-blog-api/internal/api"
    "go-blog-api/pkg/config"
)
```

### Configuration Management

- Uses Viper for YAML configuration loading
- Config file: `configs/config.yaml`
- Access via global: `config.AppConfig.Server.Port`
- Initialize in `main()` with `config.InitConfig()`

## Development Workflow

### Running the Server

```bash
# Foreground (press Ctrl+C to stop)
go run cmd/server/main.go

# Background (need to kill PID)
go run cmd/server/main.go &
```

### Finding and Stopping Background Processes

```bash
ps aux | grep "go run"  # Find PID
kill <PID>              # Stop process
lsof -i :8080           # Check port usage
```

### Testing Endpoints

```bash
curl http://localhost:8080/health
# Expected: {"status":"ok","time":"2025-12-17T22:04:00+08:00"}
```

## Project Conventions

### File Organization

- Entry point: `cmd/server/main.go` (initialization only, minimal logic)
- Private code: `internal/` (cannot be imported by external projects)
- Reusable utilities: `pkg/` (can be imported externally)
- Routes defined in: `internal/router/router.go`

### Service Constructor Pattern

Services use constructor functions returning pointers:

```go
func NewHealthService() *HealthService { return &HealthService{} }
```

### Handler Pattern

API handlers follow this structure:

```go
func HandlerName(w http.ResponseWriter, r *http.Request) {
    resp := service.NewXxxService().Method()
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(resp)
}
```

## Current State & Roadmap

- ✅ Basic HTTP server with health check endpoint
- ✅ Viper configuration loading
- ⏳ Database connection (planned: GORM with MySQL)
- ⏳ User authentication
- ⏳ Blog post CRUD operations
- ⏳ Middleware (logging, auth)
- ⏳ Unit tests

## Key Files Reference

- `cmd/server/main.go`: Application entry point
- `pkg/config/config.go`: Configuration loader and structs
- `configs/config.yaml`: Server and database configuration
- `internal/router/router.go`: Route registration
- `internal/api/health.go`: Example API handler
- `internal/service/health_service.go`: Example service implementation
