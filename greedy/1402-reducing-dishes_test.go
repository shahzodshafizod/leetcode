package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaxSatisfaction$
func TestMaxSatisfaction(t *testing.T) {
	for _, tc := range []struct {
		satisfaction []int
		maxsum       int
	}{
		{satisfaction: []int{-1, -8, 0, 5, -9}, maxsum: 14},
		{satisfaction: []int{4, 3, 2}, maxsum: 20},
		{satisfaction: []int{-1, -4, -5}, maxsum: 0},
		{satisfaction: []int{-2, 5, -1, 0, 3, -3}, maxsum: 35},
	} {
		maxsum := maxSatisfaction(tc.satisfaction)
		assert.Equal(t, tc.maxsum, maxsum)
	}
}
