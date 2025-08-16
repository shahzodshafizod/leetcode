package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestFindCrossingTime$
func TestFindCrossingTime(t *testing.T) {
	for _, tc := range []struct {
		n       int
		k       int
		time    [][]int
		minutes int
	}{
		{n: 1, k: 3, time: [][]int{{1, 1, 2, 1}, {1, 1, 3, 1}, {1, 1, 4, 1}}, minutes: 6},
		{n: 3, k: 2, time: [][]int{{1, 5, 1, 8}, {10, 10, 10, 10}}, minutes: 37},
	} {
		minutes := findCrossingTime(tc.n, tc.k, tc.time)
		assert.Equal(t, tc.minutes, minutes)
	}
}
