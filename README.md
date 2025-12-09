# 3GPP Analytics Exposure API Server

A Go server implementation of the 3GPP Analytics Exposure API (TS29522) following best coding practices.

## Overview

This project implements a RESTful API server for the 3GPP Analytics Exposure specification (TS 29.522 V16.4.0), enabling network exposure function northbound APIs for analytics data.

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go           # Server entry point
├── internal/
│   ├── handlers/
│   │   └── analytics.go       # HTTP request handlers
│   ├── models/
│   │   └── models.go          # Data models and structures
│   ├── storage/
│   │   └── store.go           # Data storage abstraction
│   └── api/
│       └── api.gen.go         # Generated OpenAPI code (if needed)
├── bin/
│   └── analytics-server       # Compiled binary
├── go.mod                      # Go module definition
├── go.sum                      # Go module checksums
├── config.yaml                 # OpenAPI code generation config
├── TS29522_AnalyticsExposure.yaml  # Main OpenAPI spec
└── README.md                   # This file
```

## Features

✅ **Full 3GPP Compliance** - Implements TS29522 Analytics Exposure API
✅ **RESTful API** - Follows REST principles with proper HTTP methods
✅ **In-Memory Storage** - Efficient temporary data storage with thread-safe operations
✅ **Comprehensive Error Handling** - Proper HTTP status codes and error messages
✅ **Middleware Support** - Logging, recovery, and CORS enabled
✅ **Health Check** - Built-in health endpoint for monitoring
✅ **OpenAPI Documentation** - Full specification available via API

## Prerequisites

- Go 1.21 or higher
- curl or similar HTTP client for testing

## Building

```bash
# Build the server binary
go build -o bin/analytics-server cmd/server/main.go

# Or use the build script
./build.sh
```

## Running the Server

```bash
# Start the server
./bin/analytics-server

# Server will start on http://localhost:8080
# API available at: http://localhost:8080/3gpp-analyticsexposure/v1
```

## API Endpoints

### Health Check

- **GET** `/health` - Server health status

### Subscriptions Management

- **GET** `/{afId}/subscriptions` - Get all subscriptions for an AF
- **POST** `/{afId}/subscriptions` - Create a new subscription
- **GET** `/{afId}/subscriptions/{subscriptionId}` - Get a specific subscription
- **PUT** `/{afId}/subscriptions/{subscriptionId}` - Update a subscription
- **DELETE** `/{afId}/subscriptions/{subscriptionId}` - Delete a subscription

### Analytics Data

- **POST** `/{afId}/fetch` - Fetch analytics data

## Testing

### Health Check

```bash
curl http://localhost:8080/health
```

### Create a Subscription

```bash
curl -X POST http://localhost:8080/3gpp-analyticsexposure/v1/AF123/subscriptions \
  -H "Content-Type: application/json" \
  -d '{
    "analyEventsSubs": [
      {
        "analyEvent": "UE_MOBILITY"
      }
    ],
    "notifUri": "http://example.com/notify",
    "notifId": "notif123"
  }'
```

### Get All Subscriptions

```bash
curl http://localhost:8080/3gpp-analyticsexposure/v1/AF123/subscriptions
```

### Fetch Analytics Data

```bash
curl -X POST http://localhost:8080/3gpp-analyticsexposure/v1/AF123/fetch \
  -H "Content-Type: application/json" \
  -d '{
    "analyEvent": "UE_MOBILITY",
    "suppFeat": "1"
  }'
```

## Architecture & Best Practices

### Design Patterns

- **Interface-based storage**: Store interface allows easy swapping of implementations
- **Handler pattern**: Clean separation of HTTP concerns from business logic
- **Middleware layer**: Proper request/response handling with logging and recovery

### Code Organization

- **cmd/**: Entry points for executables
- **internal/**: Private packages for the application
  - **handlers/**: HTTP request handlers with proper validation
  - **models/**: Data structures following 3GPP specification
  - **storage/**: Data persistence abstraction

### Thread Safety

- Uses `sync.RWMutex` for concurrent access to in-memory storage
- Safe for production use with multiple goroutines

### Error Handling

- Proper HTTP status codes (400, 404, 409, 500, etc.)
- Structured error responses
- Input validation on all endpoints

## Dependencies

- **github.com/labstack/echo/v4** - High-performance web framework
- **github.com/google/uuid** - UUID generation for subscription IDs
- **github.com/getkin/kin-openapi** - OpenAPI specification handling

## Future Enhancements

- [ ] PostgreSQL/MongoDB integration for persistent storage
- [ ] OAuth2/OpenID Connect authentication
- [ ] Rate limiting and API gateway integration
- [ ] Webhook notification delivery for subscriptions
- [ ] Metrics and observability (Prometheus)
- [ ] Docker containerization
- [ ] Kubernetes deployment manifests
- [ ] gRPC support alongside REST

## Configuration

The server runs with default configuration:

- **Port**: 8080
- **API Base Path**: `/3gpp-analyticsexposure/v1`
- **Storage**: In-memory (can be swapped)

## Monitoring

- **Health Endpoint**: `/health` - Returns JSON status
- **Logging**: Structured logs to stdout via Echo middleware
- **OpenAPI Spec**: `/openapi.json` - Full API specification

## Building from OpenAPI Spec

To regenerate code from the OpenAPI specification:

```bash
# Using oapi-codegen
oapi-codegen -config config.yaml TS29522_AnalyticsExposure.yaml > internal/api/api.gen.go
```

## References

- [3GPP TS 29.522 V16.4.0 Specification](http://www.3gpp.org/ftp/Specs/archive/29_series/29.522/)
- [OpenAPI 3.0 Specification](https://spec.openapis.org/oas/v3.0.3)
- [Echo Web Framework](https://echo.labstack.com/)

## License

This project is part of the Eurecom Multimedia Communications Lab analytics exposure project.

## Contact

For issues or contributions, please contact the development team.
