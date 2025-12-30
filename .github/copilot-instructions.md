# Copilot Instructions for go-blog-api

## Project Overview

This is a learning-focused Go web API project implementing a three-tier architecture (API/Service/Repository) for a blog backend. The project uses Gin framework with GORM for database operations.

## Architecture & Structure

### Three-Tier Pattern

- **API Layer** (`internal/api/`): HTTP handlers, request/response handling, parameter validation
- **Service Layer** (`internal/service/`): Business logic, returns `*BizError` for error handling
- **Repository Layer** (`internal/repository/`): Data access with GORM, defines interfaces for dependency injection
- **Model Layer** (`internal/model/`): GORM database models with `BaseModel` for common fields
- **DTO Layer** (`internal/dto/`): Request/Response structures, separated from models

### Key Design Patterns

```
Controller (API) → Service → Repository → Database
     ↓                ↓           ↓
   DTO            BizError    Interface
```

### Module Import Pattern

Use the module name `go-blog-api` as import prefix, group imports by: standard lib → project → third-party

```go
import (
    "strconv"

    "go-blog-api/internal/dto"
    "go-blog-api/internal/service"
    "go-blog-api/pkg/util"

    "github.com/gin-gonic/gin"
)
```

### Configuration Management

- Uses Viper for YAML configuration loading
- Config file: `configs/config.yaml`
- Includes: Server, Database, JWT settings
- Access via global: `config.AppConfig.Server.Port`, `config.AppConfig.JWT.Secret`
- Initialize in `main()` with `config.InitConfig()`

## Error Handling

### BizError Pattern

Service layer returns predefined `*BizError`, API layer uses `HandleError()`:

```go
// Service layer
if err != nil {
    return nil, util.ErrUserNotFound
}

// API layer
if err := ctrl.userService.Login(&req); err != nil {
    util.HandleError(c, err)
    return
}
```

### Predefined Errors (pkg/util/errors.go)

- `ErrInvalidParam` (40001): Parameter validation failed
- `ErrUnauthorized` (40100): Not logged in
- `ErrTokenExpired` (40101): Token expired
- `ErrForbidden` (40300): No permission
- `ErrUserNotFound` (40401): User not found
- `ErrArticleNotFound` (40402): Article not found
- `ErrUsernameExists` (40901): Username exists
- `ErrEmailExists` (40902): Email exists
- `ErrDatabase` (50001): Database error

## Development Workflow

### Running the Server

```bash
go run cmd/server/main.go
```

### Testing Endpoints

```bash
# Register
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456","email":"test@example.com"}'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'

# Get articles (with token)
curl http://localhost:8080/api/v1/articles \
  -H "Authorization: Bearer <token>"
```

## Project Conventions

### File Organization

```
cmd/server/main.go          # Entry point
configs/config.yaml         # Configuration
internal/
├── api/v1/                 # Controllers
│   ├── user.go
│   └── article.go
├── dto/                    # Request/Response DTOs
│   ├── user_dto.go
│   └── article_dto.go
├── middleware/             # Middlewares
│   └── auth.go
├── model/                  # GORM models
│   ├── base.go
│   ├── user.go
│   └── article.go
├── repository/             # Data access
│   ├── user_repo.go
│   └── article_repo.go
├── router/                 # Route registration
│   └── router.go
└── service/                # Business logic
    ├── user_service.go
    └── article_service.go
pkg/
├── config/config.go        # Config structs
├── db/db.go                # Database connection
└── util/
    ├── errors.go           # BizError definitions
    ├── response.go         # Response helpers
    └── jwt.go              # JWT utilities
```

### Controller Pattern

```go
type UserController struct {
    userService *service.UserService
}

func NewUserController() *UserController {
    repo := repository.NewUserRepository()
    svc := service.NewUserService(repo)
    return &UserController{userService: svc}
}

func (ctrl *UserController) Login(c *gin.Context) {
    var req dto.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        util.HandleError(c, util.ErrInvalidParam.WithMsg(err.Error()))
        return
    }
    
    resp, err := ctrl.userService.Login(&req)
    if err != nil {
        util.HandleError(c, err)
        return
    }
    
    util.Success(c, resp)
}
```

### Repository Interface Pattern

```go
type IUserRepository interface {
    CreateUser(user *model.User) error
    GetByUsername(username string) (*model.User, error)
}

// Compile-time interface check
var _ IUserRepository = (*UserRepository)(nil)
```

## Current State & Roadmap

- ✅ Gin HTTP server with router
- ✅ Viper configuration (server, database, JWT)
- ✅ GORM database connection with MySQL
- ✅ User authentication (register, login, JWT)
- ✅ Article CRUD operations
- ✅ JWT middleware
- ✅ Unified error handling (BizError + HandleError)
- ✅ Three-tier architecture with interfaces
- ⏳ Comment CRUD operations
- ⏳ Unit tests
- ⏳ API documentation (Swagger)

## Key Files Reference

| File | Description |
|------|-------------|
| `cmd/server/main.go` | Application entry point |
| `configs/config.yaml` | Server, database, JWT configuration |
| `pkg/config/config.go` | Configuration structs |
| `pkg/db/db.go` | Database connection |
| `pkg/util/errors.go` | BizError definitions |
| `pkg/util/response.go` | Success/Error response helpers |
| `pkg/util/jwt.go` | JWT generate/parse |
| `internal/router/router.go` | Route registration |
| `internal/middleware/auth.go` | JWT authentication middleware |
