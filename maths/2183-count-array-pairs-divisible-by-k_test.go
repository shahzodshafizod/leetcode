package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCountPairs$
func TestCountPairs(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		count int64
	}{
		{nums: []int{1, 2, 3, 4, 5}, k: 2, count: 7},
		{nums: []int{1, 2, 3, 4}, k: 5, count: 0},
	} {
		count := countPairs(tc.nums, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
