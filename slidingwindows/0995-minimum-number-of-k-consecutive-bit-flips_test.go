package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestMinKBitFlips$
func TestMinKBitFlips(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		flips int
	}{
		{nums: []int{0, 1, 0}, k: 1, flips: 2},
		{nums: []int{1, 1, 0}, k: 2, flips: -1},
		{nums: []int{0, 0, 0, 1, 0, 1, 1, 0}, k: 3, flips: 3},
		{nums: []int{0, 0, 0, 0, 0}, k: 5, flips: 1},
	} {
		flips := minKBitFlips(tc.nums, tc.k)
		assert.Equal(t, tc.flips, flips)
	}
}
