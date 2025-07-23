package predefined

import (
	"github.com/akselarzuman/containy/models"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	RedisConfig = models.Config{
		Image:        "redis:7-alpine",
		Name:         "redis-mock",
		ExposedPorts: []string{"6379:6379/tcp"},
		Strategy:     wait.ForLog("Ready to accept connections"),
	}

	PostgresConfig = func(user, password, db string) models.Config {
		return models.Config{
			Image:        "postgres:17.4-alpine",
			Name:         "postgres-mock",
			ExposedPorts: []string{"5432:5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     user,
				"POSTGRES_PASSWORD": password,
				"POSTGRES_DB":       db,
			},
			Cmd: []string{"postgres", "-c", "fsync=off"},
			Strategy: wait.ForAll(
				wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
				wait.ForListeningPort("5432/tcp"),
			),
		}
	}

	LocalstackConfig = func(services, region string) models.Config {
		return models.Config{
			Image:        "localstack/localstack:latest",
			Name:         "localstack-mock",
			ExposedPorts: []string{"4566:4566/tcp"},
			Env: map[string]string{
				"DEFAULT_REGION": region,
				"SERVICES":       services,
			},
			Strategy: wait.ForLog("Ready."),
		}
	}
)
