package models

type WaitStrategy string

const (
	WaitForLog          WaitStrategy = "log"
	WaitForPort         WaitStrategy = "port"
	WaitForHealthCheck  WaitStrategy = "health"
	WaitForHTTPResponse WaitStrategy = "http"
)
