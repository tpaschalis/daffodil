package daffodil

import (
	"net"
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

func TestHashString(t *testing.T) {
	s := "foo"
	got := hashTo16Bits(s)
	want := uint16(43507)

	assert.IsTypef(t, uint16(0), got, "Hash should be a 16-bit integer")
	assert.Equalf(t, want, got, "Hashed string should return %v, instead got %v", want, got)
}

func TestIPTo16Bits(t *testing.T) {
	ip := net.ParseIP("10.0.1.42")
	got := ipTo16Bits(ip)
	want := "foo"

	assert.IsTypef(t, uint16(0), got, "Result should be a 16-bit integer")
	assert.Equalf(t, want, got, "Transformed IP should return %v, instead got %v", want, got)

}
