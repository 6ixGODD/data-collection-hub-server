# Data Collection Hub Server
## Project Layout
```
.
├── main.go             # Entry point
├── docs                # Swagger API documentation
├── cmd                 # Command line interface
├── configs             # Configuration files
├── internal            # Source code
│   ├── app             # Application
│   └── pkg             # Internal packages
│       ├── api         # API Layer
│       ├── config      # Configuration
│       ├── dao         # Data access object
│       ├── domain      # Domain Layer
│       │   ├── entity  # Entity Struct
│       │   └── vo      # Value Object Struct
│       ├── errors      # Error handling
│       ├── middleware  # Middleware
│       ├── router      # Router Layer
│       ├── service     # Business Logic Layer
│       ├── tasks       # Scheduled tasks
│       ├── validator   # Request validation
│       └── wire        # Dependency Injection
├── pkg                 # Common packages
│   ├── cron            # Cron encapsulation
│   ├── errors          # Custom error
│   ├── jwt             # JWT encapsulation
│   ├── mongo           # MongoDB encapsulation
│   ├── prometheus      # Prometheus encapsulation
│   ├── redis           # Redis encapsulation
│   ├── utils           # Common utils
│   └── zap             # Zap encapsulation
└── test                # Test files

```

## Features
- **Web Framework**: Fiber
- **Database**: MongoDB
- **Cache**: Redis
- **Logger**: Zap
- **Configuration**: Viper
- **Dependency Injection**: Wire
- **API Documentation**: Swagger
- **Scheduler**: Cron
- **Command Line Interface**: Cobra
- **RBAC**: Casbin
- **Monitoring**: Prometheus

## Setup
```shell
make setup
```

## Reference
- https://docs.gofiber.io/
- https://github.com/goccy/go-json
- https://github.com/golang-standards/project-layout
