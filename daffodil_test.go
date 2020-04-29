package daffodil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDaffodil(t *testing.T) {
	want := &Daffodil{}
	cfg := Config{}
	got, err := NewDaffodil(cfg)

	assert.IsTypef(t, want, got, "Could not initialize Daffodil")
	assert.Nil(t, err)
}
