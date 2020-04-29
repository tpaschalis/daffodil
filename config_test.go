package daffodil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	want := &Config{}
	got, err := NewConfig()

	assert.IsTypef(t, want, got, "Could not initialize Config")
	assert.IsTypef(t, got.nodeID, uint16(0), "NodeID should be a 16-bit integer")
	assert.IsTypef(t, got.epoch, time.Now(), "Epoch should be a time.Time object")
	assert.Nil(t, err)
}
