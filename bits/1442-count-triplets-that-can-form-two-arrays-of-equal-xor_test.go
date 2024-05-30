package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestCountcount$
func TestCountTriplets(t *testing.T) {
	for _, tc := range []struct {
		arr   []int
		count int
	}{
		{arr: []int{2, 3, 1, 6, 7}, count: 4},
		{arr: []int{1, 1, 1, 1, 1}, count: 10},
	} {
		count := countTriplets(tc.arr)
		assert.Equal(t, tc.count, count)
	}
}
