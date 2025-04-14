package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCountGoodTriplets$
func TestCountGoodTriplets(t *testing.T) {
	for _, tc := range []struct {
		arr   []int
		a     int
		b     int
		c     int
		count int
	}{
		{arr: []int{3, 0, 1, 1, 9, 7}, a: 7, b: 2, c: 3, count: 4},
		{arr: []int{1, 1, 2, 2, 3}, a: 0, b: 0, c: 1, count: 0},
	} {
		count := countGoodTriplets(tc.arr, tc.a, tc.b, tc.c)
		assert.Equal(t, tc.count, count)
	}
}
