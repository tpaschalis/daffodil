package daffodil

import (
	"errors"
	"hash/fnv"
	"log"
	"net"
	"os"
	"time"
)

// Config holds configuration values for building a Daffodil.
type Config struct {
	epoch  time.Time
	nodeID uint16
}

// NewConfig initializes a Config struct
func NewConfig() (*Config, error) {
	var nodeID uint16
	var err error

	nodeIDMode := os.Getenv("DAFFODIL_NODEID_MODE")
	switch nodeIDMode {
	case "HOSTNAME":
		nodeID, err = nodeIDfromHostname()
		if err != nil {
			log.Fatal(err)
		}
	case "PRIVATEIP":
		nodeID, err = nodeIDfromIP()
		if err != nil {
			log.Fatal(err)
		}
	case "CUSTOM":
		varname := os.Getenv("DAFFODIL_NODEID_CUSTOM")
		if varname != "" {
			nodeID, err = nodeIDfromEnv(varname)
			if err != nil {
				log.Fatal(err)
			}
		}
	default:
		nodeID, err = nodeIDfromIP()
		if err != nil {
			log.Fatal(err)
		}
	}
	return &Config{
		nodeID: nodeID,
		epoch:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	}, nil
}

// This is not a cryptographically secure hash; it's simply used
// to get a numeric 16-bit 'representation' of a string,
// which could be a hostname, an environment variable or a user-provided ID.
func stringTo16Bits(s string) uint16 {
	h := fnv.New32a()
	h.Write([]byte(s))

	return uint16(h.Sum32() >> 16)
}

// IPv4 is a 32-bit address, while IPv6 is a 64-bit address.
// We obtain the final two octets of the input to provide
// a 16-bit identifier from an IP.
// IPs in the net package are stored as a byte slice of len 16.
func ipTo16Bits(ip net.IP) uint16 {
	return uint16(ip[14])<<8 + uint16(ip[15])
}

func nodeIDfromHostname() (uint16, error) {
	host, err := os.Hostname()
	if err != nil {
		return 0, err
	}

	return stringTo16Bits(host), nil
}

func nodeIDfromEnv(s string) (uint16, error) {
	val := os.Getenv(s)
	if val == "" {
		return 0, errors.New("Provided environment variable is empty")
	}

	return stringTo16Bits(val), nil
}

func nodeIDfromIP() (uint16, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return 0, errors.New("Failed to list network interfaces")
	}

	// Loop through the available network interfaces, and find the first non-loobpack.
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipTo16Bits(ipnet.IP), nil
			}
		}
	}

	return 0, errors.New("Failed to locate a valid IP address")
}
