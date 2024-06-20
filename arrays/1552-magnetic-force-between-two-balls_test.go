package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMaxDistance$
func TestMaxDistance(t *testing.T) {
	for _, tc := range []struct {
		position []int
		m        int
		distance int
	}{
		{position: []int{1, 2, 3, 4, 7}, m: 3, distance: 3},
		{position: []int{5, 4, 3, 2, 1, 1000000000}, m: 2, distance: 999999999},
	} {
		distance := maxDistance(tc.position, tc.m)
		assert.Equal(t, tc.distance, distance)
	}
}
