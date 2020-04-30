package daffodil

// Define bit lengths of parts consisting.
const (
	Time   = 39
	Order  = 8
	NodeID = 63 - Time - Order // 16-bit

	daffodilTimeUnit = 1e7 // == 10 msec

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
