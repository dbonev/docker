package builder

import "github.com/docker/docker/runconfig"

// Image represents a Docker image used by the builder.
type Image interface {
	ID() string
	Config() *runconfig.Config
}
