package models

type Config struct {
	Image        string
	Name         string
	ExposedPorts []string
	Env          map[string]string
	Cmd          []string

	WaitStrategy WaitStrategy
	WaitConfig   map[string]string
}
