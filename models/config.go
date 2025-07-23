package models

import "github.com/testcontainers/testcontainers-go/wait"

type Config struct {
	Image        string
	Name         string
	ExposedPorts []string
	Env          map[string]string
	Cmd          []string
	Strategy     wait.Strategy
}
