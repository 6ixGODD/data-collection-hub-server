# Data Collection Hub Server
## Description
Directory structure of the project
```shell
.
├── README.md         # Project README
├── api               # API definition
│    └── v1  
│        ├── admin  
│        ├── common
│        └── user
├── cmd               # Command line interface
├── core              # Core module
│    ├── config       # Configuration module
│    │     ├── config.go
│    │     └── modules
│    ├── memcached    # Memcached module
│    ├── mongo        # MongoDB module
│    ├── redis        # Redis module
│    ├── viper        # Viper module
│    ├── wire         # Wire module
│    └── zap          # Zap logger module
├── dal               # Data access object
├── docs              # Swagger API documentation
├── global            # Global variables
├── go.mod            # Go module file
├── go.sum 
├── initializer       # Initializer and application factory
├── main.go           # Main entry
├── middleware        # Middleware
│    ├── cors.go      # CORS middleware
│    ├── jwt.go       # JWT middleware
│    ├── limit_ip.go  # IP limit middleware
│    └── logger.go    # Logger middleware
├── models            # Data models
│    ├── schema       # Request and response schema
│    │     ├── admin
│    │     ├── common
│    │     ├── response.go
│    │     └── user
├── router            # Router
├── scheduler         # Scheduler and timer
├── service           # Service layer
│    ├── admin
│    ├── common
│    └── user
└── utils             # Utilities
    ├── check         # Check utilities
    ├── crypt         # Cryptography
    ├── ip            # IP utilities
    ├── jwt           # JWT utilities
    └── validate      # Validate utilities
```

## Technology Stack
- **Framework**: Fiber
- **DB**: MongoDB
- **Cache**: Redis, GoMemcached
- **Logger**: Zap
- **Configuration**: Viper
- **Dependency Injection**: Wire
- **API Documentation**: Swagger
- **Scheduler**: Cron
- **Command Line Interface**: Cobra
- **Middleware**:
  - CORS
  - JWT
  - IP limit
  - Logger
  - Casbin
- **RBAC**: Casbin

## Setup
```shell
go mod tidy
```

## Reference
- https://docs.gofiber.io/
- https://github.com/goccy/go-json