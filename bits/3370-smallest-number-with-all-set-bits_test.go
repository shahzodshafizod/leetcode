package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestSmallestNumber$
func TestSmallestNumber(t *testing.T) {
	for _, tc := range []struct {
		n        int
		smallest int
	}{
		{n: 5, smallest: 7},
		{n: 10, smallest: 15},
		{n: 3, smallest: 3},
	} {
		smallest := smallestNumber(tc.n)
		assert.Equal(t, tc.smallest, smallest)
	}
}
