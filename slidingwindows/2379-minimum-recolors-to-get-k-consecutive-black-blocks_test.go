package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestMinimumRecolors$
func TestMinimumRecolors(t *testing.T) {
	for _, tc := range []struct {
		blocks string
		k      int
		ops    int
	}{
		{blocks: "WBBWWBBWBW", k: 7, ops: 3},
		{blocks: "WBWBBBW", k: 2, ops: 0},
		{blocks: "BWWWBB", k: 6, ops: 3},
	} {
		ops := minimumRecolors(tc.blocks, tc.k)
		assert.Equal(t, tc.ops, ops)
	}
}
