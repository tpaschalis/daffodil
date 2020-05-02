package daffodil

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDaffodil(t *testing.T) {
	want := &Daffodil{}
	cfg := &Config{}
	got, err := NewDaffodil(cfg)

	assert.IsTypef(t, want, got, "Could not initialize Daffodil")
	assert.IsTypef(t, cfg, got.cfg, "Daffodil doesn't contain Config")
	assert.Nil(t, err)
}

func TestNextMethod(t *testing.T) {
	cfg, err := NewConfig()
	require.Nil(t, err)
	d, err := NewDaffodil(cfg)
	require.Nil(t, err)

	got, err := d.Next()
	assert.Nil(t, err)
	assert.IsType(t, ID(0), got, "Next method should return an ID")
	assert.NotEmpty(t, got)

	// Debugging printouts, it seems to work!!
	// d.sequence = 100
	// fmt.Println("timestamp", timestamp, uint64(timestamp))
	// fmt.Println(uint64(timestamp) << (NodeBits + SequenceBits))
	// fmt.Println(strconv.FormatUint(uint64(timestamp)<<NodeBits<<SequenceBits, 2))
	// fmt.Println("")

	// fmt.Println("nodeID", d.cfg.nodeID)
	// fmt.Println(uint64(d.cfg.nodeID) << SequenceBits)
	// fmt.Println(strconv.FormatUint(uint64(d.cfg.nodeID)<<SequenceBits, 2))
	// fmt.Println("")

	// fmt.Println("Sequence")
	// fmt.Println(uint64(d.sequence))
	// fmt.Println(strconv.FormatUint(uint64(d.sequence), 2))
	// fmt.Println("")

}

func TestGetTicks(t *testing.T) {
	cfg, err := NewConfig()
	require.Nil(t, err)
	d, err := NewDaffodil(cfg)
	require.Nil(t, err)
	got := d.getTicks()
	assert.NotEmpty(t, got)
}

func TestPrintMasks(t *testing.T) {
	fmt.Println(strconv.FormatInt(timeMask, 2))
	fmt.Println(strconv.FormatInt(nodeMask, 2))
	fmt.Println(strconv.FormatInt(seqMask, 2))
	// timestamp --> node-id --> sequence
	// 111111111111111111111111111111111111111000000000000000000000000
	//                                        111111111111111100000000
	//                                                        11111111
}

func TestDismantleID(t *testing.T) {
	id := ID(17801782472864612)

	gotTs, gotNode, gotSeq := id.dismantle()
	assert.Equal(t, int64(1061068920), gotTs)
	assert.Equal(t, int64(43507), gotNode)
	assert.Equal(t, int64(100), gotSeq)
}
