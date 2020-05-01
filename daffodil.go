package daffodil

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
	cfg Config
}

// ID represents the generated 64-bit UID.
type ID uint64

// NewDaffodil returns a new instance of an ID generator.
func NewDaffodil(cfg Config) (*Daffodil, error) {
	return &Daffodil{}, nil
}

// Next generates the next uid.
func (d *Daffodil) Next() ID {
	return 0
}
