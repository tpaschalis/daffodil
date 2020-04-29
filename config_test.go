package daffodil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	want := &Config{}
	got, err := NewConfig()

	assert.IsTypef(t, want, got, "Could not initialize Config")
	assert.IsTypef(t, got.nodeID, uint16(0), "NodeID should be a 16-bit integer")
	assert.Nil(t, err)
}
