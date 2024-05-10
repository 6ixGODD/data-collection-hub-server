# Data Collection Hub Server
## Project Layout
```
.
├── Makefile                # Makefile for build, run, test, etc.
├── main.go                 # Entry point of the application
├── api                     # API documentation
├── cmd                     # Command line interface based on Cobra
│   └── root.go             # Root command
├── configs                 # Configuration files
│   ├── dev                 # Development environment
│   ├── test                # Test environment
│   └── prob                # Production environment
├── internal                # Internal packages
│   ├── app                 # Application Factory
│   └── pkg                 # Internal packages
│       ├── api             # API Layer
│       ├── config          # Configuration
│       ├── dal             # Data Access Layer
│       ├── hooks           # Hooks
│       ├── models          # Data Models
│       ├── router          # Router Layer
│       ├── scheduler       # Scheduler Tasks
│       ├── schema          # Request and Response Schema
│       └── service         # Business Logic Layer
├── pkg                     # Common packages
│   ├── casbin              # Casbin RBAC
│   ├── cron                # Cron Scheduler
│   ├── middleware          # Middleware
│   ├── mongo               # MongoDB
│   ├── prometheus          # Prometheus
│   ├── redis               # Redis
│   ├── utils               # Utilities
│   ├── wire                # Dependency Injection
│   └── zap                 # Logger
└── test                    # Test files
    ├── api                 # API Test
    ├── dal                 # DAL Test
    ├── service             # Service Test
    └── utils               # Utilities Test

```

## Features
- **Framework**: Fiber
- **DB**: MongoDB
- **Cache**: Redis
- **Logger**: Zap
- **Configuration**: Viper
- **Dependency Injection**: Wire
- **API Documentation**: Swagger
- **Scheduler**: Cron
- **Command Line Interface**: Cobra
- **RBAC**: Casbin
- **Monitoring**: Prometheus
- **Middleware**:
  - CORS
  - JWT
  - IP limit
  - Logger
  - Casbin
  - Prometheus

## Setup
```shell
go mod tidy
```

## Reference
- https://docs.gofiber.io/
- https://github.com/goccy/go-json
- https://github.com/golang-standards/project-layout
