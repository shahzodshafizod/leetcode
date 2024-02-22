package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestReverseBits$
func TestReverseBits(t *testing.T) {
	for _, tc := range []struct {
		num      uint32
		reversed uint32
	}{
		{num: 0b00000010100101000001111010011100, reversed: 964176192},
		{num: 0b11111111111111111111111111111101, reversed: 3221225471},
	} {
		reversed := reverseBits(tc.num)
		assert.Equal(t, tc.reversed, reversed)
	}
}
