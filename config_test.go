package daffodil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	want := &Config{}
	got, err := NewConfig()

	assert.IsTypef(t, got.nodeID, uint16(0), "NodeID should be a 16-bit integer")
	assert.IsTypef(t, got.epoch, time.Now(), "Epoch should be a time.Time object")
	assert.IsTypef(t, want, got, "Could not initialize Config")
	assert.Nil(t, err)
}

func TestHashString(t *testing.T) {
	s := "foo"
	got := hashTo16Bits(s)
	want := uint16(43507)

	assert.IsTypef(t, got, uint16(0), "Hash should be a 16-bit integer")
	assert.Equalf(t, want, got, "Hashed string should return %v, instead got %v", got, want)
}
