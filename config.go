package daffodil

// Config holds configuration values for building a Daffodil
type Config struct {
	nodeID uint16
}

// NewConfig initializes a Config struct
func NewConfig() (*Config, error) {
	return &Config{
		nodeID: 0,
	}, nil
}
