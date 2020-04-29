package daffodil

// Config holds configuration values for building a Daffodil
type Config struct{}

// NewConfig initializes a Config struct
func NewConfig() (*Config, error) {
	return &Config{}, nil
}
