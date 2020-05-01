package daffodil

import (
	"fmt"
	"sync"
	"time"
)

// Define bit lengths of parts consisting.
const (
	TimeBits     = 39
	SequenceBits = 8
	NodeBits     = 63 - TimeBits - SequenceBits // 16-bit

	daffodilTimeUnit = 1e7 // == 10 msec

	nodeMask  = (1<<NodeBits - 1)
	orderMask = (1<<SequenceBits - 1) << NodeBits
	timeMask  = (1<<TimeBits - 1) << SequenceBits << NodeBits
)

// Daffodil is an id generator.
type Daffodil struct {
	cfg     *Config
	mutex   *sync.Mutex
	elapsed time.Duration
}

// ID represents the generated 64-bit UID.
type ID uint64

// NewDaffodil returns a new instance of an ID generator.
func NewDaffodil(cfg *Config) (*Daffodil, error) {
	return &Daffodil{
		cfg: cfg,
	}, nil
}

// Next generates the next uid.
func (d *Daffodil) Next() ID {
	return 0
}

func (d *Daffodil) getTicks() int64 {
	fmt.Println(time.Now())
	fmt.Println(d.cfg.epoch)
	return time.Now().UTC().UnixNano()/daffodilTimeUnit -
		d.cfg.epoch.UTC().UnixNano()/daffodilTimeUnit
}
