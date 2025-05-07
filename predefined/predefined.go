package predefined

import "github.com/akselarzuman/containy/models"

var (
	RedisConfig = models.Config{
		Image:        "redis:7-alpine",
		Name:         "redis-mock",
		ExposedPorts: []string{"6379:6379/tcp"},
		WaitStrategy: models.WaitForLog,
		WaitConfig: map[string]string{
			"log": "Ready to accept connections",
		},
	}

	PostgresConfig = func(user, password, db string) models.Config {
		return models.Config{
			Image:        "postgres:17.2-alpine",
			Name:         "postgres-mock",
			ExposedPorts: []string{"5432:5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     user,
				"POSTGRES_PASSWORD": password,
				"POSTGRES_DB":       db,
			},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			WaitStrategy: models.WaitForLog,
			WaitConfig: map[string]string{
				"log": "database system is ready to accept connections",
			},
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
			WaitStrategy: models.WaitForLog,
			WaitConfig: map[string]string{
				"log": "Ready.",
			},
		}
	}
)
