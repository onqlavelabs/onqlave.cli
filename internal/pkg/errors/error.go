package errors

import (
	"fmt"
)

type Error struct {
	service string
	message string
	baseErr errorI
	mode    Mode
}

type errorConfig struct {
	service string
	mode    Mode
}

var errConfig = errorConfig{}

func InitErrorConfig(serviceName string, mode string) {
	m := 0
	if mode == "DEBUG" {
		m = 1
	}

	errConfig = errorConfig{
		service: serviceName,
		mode:    Mode(m),
	}
}

type errorI interface {
	Error() string
}

func (err Error) Error() string {
	switch err.mode {
	case DEBUG:
		if err.baseErr != nil {
			return fmt.Sprintf("Service %s encounter an error : %s -- base error: %s", err.service, err.message, err.baseErr.Error())
		}
	}

	return fmt.Sprintf("Service %s encounter an error : %s", err.service, err.message)
}
