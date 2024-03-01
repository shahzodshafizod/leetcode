package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestMaximumOddBinaryNumber$
func TestMaximumOddBinaryNumber(t *testing.T) {
	for _, tc := range []struct {
		s      string
		maxOdd string
	}{
		{s: "010", maxOdd: "001"},
		{s: "0101", maxOdd: "1001"},
		{s: "1", maxOdd: "1"},
		{s: "100000000", maxOdd: "000000001"},
	} {
		maxOdd := maximumOddBinaryNumber(tc.s)
		assert.Equal(t, tc.maxOdd, maxOdd)
	}
}
