package models

import "github.com/testcontainers/testcontainers-go/wait"

type Config struct {
	Image        string
	Name         string
	ExposedPorts []string
	PortBindings map[string]string // container port (e.g. "4566/tcp") -> host port (e.g. "4566")
	Env          map[string]string
	Cmd          []string
	Strategy     wait.Strategy
}
