# containy

`containy` is a Go library that simplifies container-based testing by providing a clean, declarative API on top of [testcontainers-go](https://github.com/testcontainers/testcontainers-go). It helps you write more maintainable integration tests with less boilerplate.

## Features

- Declarative container configuration
- Flexible wait strategies (logs, ports, health checks, HTTP)
- [Predefined](https://github.com/akselarzuman/containy/blob/main/predefined/predefined.go) templates for common services (Redis, PostgreSQL, Localstack)
- Easy to extend and customize

## Installation

```bash
go get github.com/akselarzuman/containy
```

## Quick Start
See [examples](https://github.com/akselarzuman/containy/blob/main/examples/main.go) for a complete example.

## Predefined Services

### Redis
```go
redis, err := c.CreateContainer(ctx, predefined.RedisConfig)
```

### PostgreSQL
```go
postgres, err := c.CreateContainer(ctx, predefined.PostgresConfig(
    "postgres", 
    "password",
    "testdb"))
```

### Localstack
```go
localstack, err := c.CreateContainer(ctx, predefined.LocalstackConfig(
    "dynamodb,s3",
    "us-east-1",
))
```

## Custom Container Configuration

Create custom configurations for any container:

```go
config := models.Config{
    Image:        "nginx:latest",
    Name:         "nginx-test",
    ExposedPorts: []string{"80:80/tcp"},
    WaitStrategy: models.WaitForHTTPResponse,
    WaitConfig: map[string]string{
        "path": "/",
        "port": "80/tcp",
    },
}

container, err := c.CreateContainer(ctx, config)
```
Custom postgres configuration:
```go
config := models.Config{
    Image:        "postgres:13",
    Name:         "postgres-mock",
    ExposedPorts: []string{"5555:5432/tcp"},
    Env: map[string]string{
        "POSTGRES_USER":     "postgres",
        "POSTGRES_PASSWORD": "pgpwd",
        "POSTGRES_DB":       "containy",
    },
    Cmd:          []string{"postgres", "-c", "fsync=off"},
    WaitStrategy: models.WaitForPort,
    WaitConfig: map[string]string{
        "port": "5432/tcp",
    },
}

container, err := c.CreateContainer(ctx, config)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.