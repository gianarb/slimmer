package runner

import (
	"io"
)

type Runner interface {
	BuildContainer(img string, envVars []string, volumes []string, workDir string) (string, error)
	Exec(containerId string, command []string) (int, error)
	RemoveContainer(containerId string) error
	GetStream() io.Writer
}
