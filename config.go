package daffodil

import (
	"errors"
	"hash/fnv"
	"net"
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

// IPv4 is a 32-bit address, while IPv6 is a 64-bit address
// We obtain the final two octets of the input to provide
// a 16-bit identifier from an IP.
// IPs in the net package are stored as a byte slice of len 16
func ipTo16Bits(ip net.IP) uint16 {
	return uint16(ip[14])<<8 + uint16(ip[15])
}

func nodeIDfromHostname() (uint16, error) {
	host, err := os.Hostname()
	if err != nil {
		return 0, err
	}

	return hashTo16Bits(host), nil
}

func nodeIDfromEnv(s string) (uint16, error) {
	val := os.Getenv(s)
	if val == "" {
		return 0, errors.New("Provided environment variable is empty")
	}

	return hashTo16Bits(val), nil
}
