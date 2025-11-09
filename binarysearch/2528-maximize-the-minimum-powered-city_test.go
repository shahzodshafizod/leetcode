package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestMaxPower$
func TestMaxPower(t *testing.T) {
	for _, tc := range []struct {
		stations []int
		r        int
		k        int
		power    int64
	}{
		{stations: []int{1, 2, 4, 5, 0}, r: 1, k: 2, power: 5},
		{stations: []int{4, 4, 4, 4}, r: 0, k: 3, power: 4},
		{stations: []int{2, 10, 12, 3}, r: 0, k: 14, power: 9},
	} {
		power := maxPower(tc.stations, tc.r, tc.k)
		assert.Equal(t, tc.power, power)
	}
}
