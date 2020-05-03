package daffodil

import (
	"errors"
	"sync"
	"time"
)

// Define bit lengths of ID compomnents.
const (
	TimeBits     = 39
	SequenceBits = 8
	NodeBits     = 63 - TimeBits - SequenceBits // 16-bit

	daffodilTimeUnit = 1e7 // == 10 msec

	seqMask  = (1<<SequenceBits - 1)
	nodeMask = (1<<NodeBits - 1) << SequenceBits
	timeMask = (1<<TimeBits - 1) << NodeBits << SequenceBits
)

// Daffodil is an id generator.
type Daffodil struct {
	cfg      *Config
	mutex    *sync.Mutex
	latest   int64
	sequence int64
}

// ID represents the generated 64-bit UID.
type ID uint64

// NewDaffodil returns a new instance of an ID generator.
func NewDaffodil(cfg *Config) (*Daffodil, error) {
	return &Daffodil{
		cfg:   cfg,
		mutex: new(sync.Mutex),
	}, nil
}

// Next generates the next uid.
func (d *Daffodil) Next() (ID, error) {

	ticks := d.getTicks()
	timestamp := ticks

	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Has clock moved backwards or wrapped around?
	if timestamp < d.latest || ticks < 0 {
		return 0, errors.New("clock has moved backwards or maybe has wrapped around")
	}

	// Check if we've run out of sequence bits in this tick, otherwise increment sequence
	if timestamp == d.latest {
		if d.sequence >= seqMask {
			return 0, errors.New("sequence overflow, not generating IDs for the rest of the tick")
		}
		d.sequence++
	} else {
		// we're in a new-er timeslot, reset the sequence and store the next timestamp
		d.sequence = 0
		d.latest = timestamp
	}

	return ID(uint64(timestamp)<<(NodeBits+SequenceBits) |
		uint64(d.cfg.nodeID)<<SequenceBits |
		uint64(d.sequence)), nil
}

func (d *Daffodil) getTicks() int64 {
	return time.Now().UTC().UnixNano()/daffodilTimeUnit -
		d.cfg.epoch.UTC().UnixNano()/daffodilTimeUnit
}

// Dismantle breaks up an ID to its basic compomnents.
func (id ID) Dismantle() (int64, int64, int64) {
	return int64(id & timeMask >> NodeBits >> SequenceBits),
		int64(id & nodeMask >> SequenceBits),
		int64(id & seqMask)
}
