package bits

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestGetSneakyNumbers$
func TestGetSneakyNumbers(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		sneaky []int
	}{
		{nums: []int{0, 1, 1, 0}, sneaky: []int{0, 1}},
		{nums: []int{0, 3, 2, 1, 3, 2}, sneaky: []int{2, 3}},
		{nums: []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}, sneaky: []int{4, 5}},
	} {
		sneaky := getSneakyNumbers(tc.nums)
		sort.Ints(sneaky)
		sort.Ints(tc.sneaky)
		assert.Equal(t, tc.sneaky, sneaky)
	}
}
