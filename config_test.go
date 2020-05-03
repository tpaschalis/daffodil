package daffodil

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	want := &Config{}
	got, err := NewConfig()

	assert.IsTypef(t, uint16(0), got.nodeID, "NodeID should be a 16-bit integer")
	assert.IsTypef(t, time.Now(), got.epoch, "Epoch should be a time.Time object")
	assert.IsTypef(t, want, got, "Could not initialize Config")
	assert.Nil(t, err)
}

func TestStringTo16Bits(t *testing.T) {
	s := "foo"
	got := stringTo16Bits(s)
	want := uint16(43507)

	assert.IsTypef(t, uint16(0), got, "Hash should be a 16-bit integer")
	assert.Equalf(t, want, got, "Hashed string should return %v, instead got %v", want, got)
}

func TestIPTo16Bits(t *testing.T) {
	// The final two octets are '1' and '42'.
	// We expect the result to be 1 * 2^8 + 42 = 298
	ip := net.ParseIP("10.0.1.42")
	got := ipTo16Bits(ip)
	want := uint16(298)

	assert.IsTypef(t, uint16(0), got, "Result should be a 16-bit integer")
	assert.Equalf(t, want, got, "Transformed IP should return %v, instead got %v", want, got)

}

func TestNodeIDSelection(t *testing.T) {

	// Setting up NodeID via hostname
	os.Setenv("DAFFODIL_NODEID_MODE", "HOSTNAME")

	cfg1, err := NewConfig()
	assert.Nil(t, err)

	nid1, err := nodeIDfromHostname()
	assert.Nil(t, err)
	assert.Equal(t, cfg1.nodeID, nid1)

	// Setting up NodeID via custom env var
	os.Setenv("DAFFODIL_NODEID_MODE", "CUSTOM")
	os.Setenv("DAFFODIL_NODEID_CUSTOM", "DFDL_NID")
	os.Setenv("DFDL_NID", "bar")

	cfg2, err := NewConfig()
	assert.Nil(t, err)
	nid2, err := nodeIDfromEnv("DFDL_NID")
	assert.Nil(t, err)

	assert.Equal(t, cfg2.nodeID, nid2) // bar --> 30391
}
