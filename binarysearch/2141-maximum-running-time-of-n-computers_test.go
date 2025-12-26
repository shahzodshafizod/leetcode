package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestMaxRunTime$
func TestMaxRunTime(t *testing.T) {
	for _, tc := range []struct {
		n         int
		batteries []int
		time      int64
	}{
		{n: 2, batteries: []int{3, 3, 3}, time: 4},
		{n: 2, batteries: []int{1, 1, 1, 1}, time: 2},
		{n: 1, batteries: []int{1, 1, 1}, time: 3},
		{n: 3, batteries: []int{10, 10, 3, 5}, time: 8},
		{n: 2, batteries: []int{1, 2, 3, 4, 5}, time: 7},
	} {
		time := maxRunTime(tc.n, tc.batteries)
		assert.Equal(t, tc.time, time)
	}
}
