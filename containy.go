package containy

import (
	"context"
	"errors"
	"fmt"

	"github.com/akselarzuman/containy/models"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type containy struct {
	containers []testcontainers.Container
}

func New() *containy {
	return &containy{}
}

func (c *containy) CreateContainer(ctx context.Context, config models.Config) (testcontainers.Container, error) {
	waitStrategy, err := c.getWaitStrategy(config)
	if err != nil {
		return nil, fmt.Errorf("failed to configure wait strategy: %w", err)
	}

	req := testcontainers.ContainerRequest{
		Image:        config.Image,
		Name:         config.Name,
		ExposedPorts: config.ExposedPorts,
		Env:          config.Env,
		Cmd:          config.Cmd,
		WaitingFor:   waitStrategy,
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create container: %w", err)
	}

	c.containers = append(c.containers, container)
	return container, nil
}

func (c *containy) Cleanup(ctx context.Context) error {
	var errs []error
	for _, container := range c.containers {
		if err := container.Terminate(ctx); err != nil {
			errs = append(errs, fmt.Errorf("failed to terminate container %s: %w", container.GetContainerID(), err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to cleanup containers: %v", errs)
	}

	return nil
}

func (containy) getWaitStrategy(config models.Config) (wait.Strategy, error) {
	switch config.WaitStrategy {
	case models.WaitForLog:
		logStr, ok := config.WaitConfig["log"]
		if !ok {
			return nil, errors.New("log wait strategy requires 'log' in WaitConfig")
		}

		return wait.ForLog(logStr), nil
	case models.WaitForPort:
		port, ok := config.WaitConfig["port"]
		if !ok {
			return nil, errors.New("port wait strategy requires 'port' in WaitConfig")
		}

		return wait.ForListeningPort(nat.Port(port)), nil
	case models.WaitForHealthCheck:
		return wait.ForHealthCheck(), nil
	case models.WaitForHTTPResponse:
		path, ok := config.WaitConfig["path"]
		if !ok {
			return nil, errors.New("HTTP wait strategy requires 'path' in WaitConfig")
		}

		port, ok := config.WaitConfig["port"]
		if !ok {
			return nil, errors.New("HTTP wait strategy requires 'port' in WaitConfig")
		}

		return wait.ForHTTP(path).WithPort(nat.Port(port)), nil
	default:
		return nil, fmt.Errorf("unknown wait strategy: %s", config.WaitStrategy)
	}
}
