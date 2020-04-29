package daffodil

// Define bit lengths of parts consisting
const (
	Time   = 39
	Order  = 8
	NodeID = 63 - Time - Order // 16-bit
)

// Daffodil is an id generator
type Daffodil struct{}

// NewDaffodil returns a new instance of an ID generator
func NewDaffodil(cfg Config) (*Daffodil, error) {
	return &Daffodil{}, nil
}
