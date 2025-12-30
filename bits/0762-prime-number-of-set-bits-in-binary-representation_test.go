package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestCountPrimeSetBits$
func TestCountPrimeSetBits(t *testing.T) {
	for _, tc := range []struct {
		left   int
		right  int
		output int
	}{
		{left: 6, right: 10, output: 4},
		{left: 10, right: 15, output: 5},
		{left: 1, right: 1, output: 0},  // 1 has 1 set bit, 1 is not prime
		{left: 2, right: 2, output: 0},  // 2 has 1 set bit, 1 is not prime
		{left: 3, right: 3, output: 1},  // 3 has 2 set bits, 2 is prime
		{left: 4, right: 4, output: 0},  // 4 has 1 set bit, 1 is not prime
		{left: 5, right: 5, output: 1},  // 5 has 2 set bits, 2 is prime
		{left: 1, right: 10, output: 6}, // 3,5,6,7,9,10 have prime set bits
		{left: 244, right: 269, output: 16},
	} {
		output := countPrimeSetBits(tc.left, tc.right)
		assert.Equal(t, tc.output, output)
	}
}
