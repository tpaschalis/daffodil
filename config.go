package daffodil

import (
	"errors"
	"hash/fnv"
	"os"
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

// This is not a cryptographically secure hash; it's simply used
// to get a numeric 16-bit representation of a string, be it a
// hostname, an environment variable or a user-provided ID.
func hashTo16Bits(s string) uint16 {
	h := fnv.New32a()
	h.Write([]byte(s))

	return uint16(h.Sum32() >> 16)
}

func getIDfromHostname() (uint16, error) {
	host, err := os.Hostname()
	if err != nil {
		return 0, err
	}

	return hashTo16Bits(host), nil
}

func getIDfromEnv(s string) (uint16, error) {
	val := os.Getenv(s)
	if val == "" {
		return 0, errors.New("Provided environment variable is empty")
	}

	return hashTo16Bits(val), nil
}
