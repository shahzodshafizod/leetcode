package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestMinSwapsCouples$
func TestMinSwapsCouples(t *testing.T) {
	for _, tc := range []struct {
		row   []int
		swaps int
	}{
		{row: []int{0, 2, 1, 3}, swaps: 1},
		{row: []int{3, 2, 0, 1}, swaps: 0},
		{row: []int{3, 0, 2, 1}, swaps: 1},
		{row: []int{0, 2, 3, 4, 5, 1, 6, 8, 7, 9}, swaps: 3},
	} {
		swaps := minSwapsCouples(tc.row)
		assert.Equal(t, tc.swaps, swaps)
	}
}
