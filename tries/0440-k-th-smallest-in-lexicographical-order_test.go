package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestFindKthNumber$
func TestFindKthNumber(t *testing.T) {
	for _, tc := range []struct {
		n   int
		k   int
		kth int
	}{
		{n: 13, k: 2, kth: 10},
		{n: 1, k: 1, kth: 1},
		{n: 1000000000, k: 1000000000, kth: 999999999},
		{n: 957747794, k: 424238336, kth: 481814499},
	} {
		kth := findKthNumber(tc.n, tc.k)
		assert.Equal(t, tc.kth, kth)
	}
}
