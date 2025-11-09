package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinTaps$
func TestMinTaps(t *testing.T) {
	for _, tc := range []struct {
		n      int
		ranges []int
		taps   int
	}{
		{n: 5, ranges: []int{3, 4, 1, 1, 0, 0}, taps: 1},
		{n: 3, ranges: []int{0, 0, 0, 0}, taps: -1},
	} {
		taps := minTaps(tc.n, tc.ranges)
		assert.Equal(t, tc.taps, taps)
	}
}
