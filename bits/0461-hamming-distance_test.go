package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestHammingDistance$
func TestHammingDistance(t *testing.T) {
	for _, tc := range []struct {
		start int
		goal  int
		flips int
	}{
		{start: 10, goal: 7, flips: 3},
		{start: 3, goal: 4, flips: 3},
	} {
		flips := hammingDistance(tc.start, tc.goal)
		assert.Equal(t, tc.flips, flips)
	}
}
