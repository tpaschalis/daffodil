package daffodil

// Daffodil is an id generator
type Daffodil struct{}

// NewDaffodil returns a new instance of an ID generator
func NewDaffodil(cfg Config) (*Daffodil, error) {
	return &Daffodil{}, nil
}
