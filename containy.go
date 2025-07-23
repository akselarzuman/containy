package containy

import (
	"context"
	"fmt"

	"github.com/akselarzuman/containy/models"
	"github.com/testcontainers/testcontainers-go"
)

type Containy struct {
	containers []testcontainers.Container
}

func New() *Containy {
	return &Containy{}
}

func (c *Containy) CreateContainer(ctx context.Context, config models.Config) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        config.Image,
		Name:         config.Name,
		ExposedPorts: config.ExposedPorts,
		Env:          config.Env,
		Cmd:          config.Cmd,
		WaitingFor:   config.Strategy,
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

func (c *Containy) Cleanup(ctx context.Context) error {
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
