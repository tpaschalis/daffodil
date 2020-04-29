package daffodil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	want := &Config{}
	got, err := NewConfig()

	assert.IsTypef(t, want, got, "Could not initialize Config")
	assert.Nil(t, err)
}
