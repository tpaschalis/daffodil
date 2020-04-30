package daffodil

import (
	"time"
)

// Config holds configuration values for building a Daffodil
type Config struct {
	epoch  time.Time
	nodeID uint16
}

// NewConfig initializes a Config struct
func NewConfig() (*Config, error) {
	return &Config{
		nodeID: 0,
		epoch:  time.Now(),
	}, nil
}
