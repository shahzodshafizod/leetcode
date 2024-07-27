package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMinimumCost$
func TestMinimumCost(t *testing.T) {
	for _, tc := range []struct {
		source   string
		target   string
		original []byte
		changed  []byte
		cost     []int
		minCost  int64
	}{
		{
			source:   "abcd",
			target:   "acbe",
			original: []byte{'a', 'b', 'c', 'c', 'e', 'd'},
			changed:  []byte{'b', 'c', 'b', 'e', 'b', 'e'},
			cost:     []int{2, 5, 5, 1, 2, 20},
			minCost:  28,
		},
		{
			source:   "aaaa",
			target:   "bbbb",
			original: []byte{'a', 'c'},
			changed:  []byte{'c', 'b'},
			cost:     []int{1, 2},
			minCost:  12,
		},
		{
			source:   "abcd",
			target:   "abce",
			original: []byte{'a'},
			changed:  []byte{'e'},
			cost:     []int{10000},
			minCost:  -1,
		},
		{
			source:   "bcaabaddac",
			target:   "bdccbdaadc",
			original: []byte{'c', 'd', 'a', 'a', 'c', 'a', 'd'},
			changed:  []byte{'a', 'a', 'd', 'b', 'd', 'c', 'c'},
			cost:     []int{4, 3, 6, 3, 11, 6, 4},
			minCost:  40,
		},
	} {
		minCost := minimumCost(tc.source, tc.target, tc.original, tc.changed, tc.cost)
		assert.Equal(t, tc.minCost, minCost)
	}
}
