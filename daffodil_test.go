package daffodil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDaffodil(t *testing.T) {
	want := &Daffodil{}
	cfg := Config{}
	got, err := NewDaffodil(cfg)

	assert.IsTypef(t, want, got, "Could not initialize Daffodil")
	assert.IsTypef(t, cfg, got.cfg, "Daffodil doesn't contain Config")
	assert.Nil(t, err)
}

func TestNextMethod(t *testing.T) {
	want := ID(0)
	d, err := NewDaffodil(Config{})
	require.Nil(t, err)

	got := d.Next()
	assert.IsType(t, want, got, "Next method should return an ID")
}

func TestGetTicks(t *testing.T) {
	d, err := NewDaffodil(Config{})
	d.elapsed = time.Now().Sub(d.cfg.epoch)

	got := d.getTicks()
	assert.NotEmpty(t, got)
}
