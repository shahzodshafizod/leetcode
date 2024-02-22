package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestHammingWeight$
func TestHammingWeight(t *testing.T) {
	for _, tc := range []struct {
		num  uint32
		ones int
	}{
		{num: 0b00000000000000000000000000001011, ones: 3},
		{num: 0b00000000000000000000000010000000, ones: 1},
		{num: 0b11111111111111111111111111111101, ones: 31},
	} {
		ones := hammingWeight(tc.num)
		assert.Equal(t, tc.ones, ones)
	}
}
