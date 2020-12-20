package pluginamzn

import (
	"context"

	"github.com/cgentron/api/iface"
)

// Config ...
type Config struct{}

// CreateConfig ...
func CreateConfig() *Config {
	return &Config{}
}

// Amazn ...
type Amazn struct {
	name string
}

// New ...
func New(config *Config, name string) (iface.ResolverHandler, error) {
	return &Amazn{
		name: name,
	}, nil
}

func (a *Amazn) Resolve(context context.Context, msg []byte) ([]byte, error) {
	return []byte("amazn"), nil
}
