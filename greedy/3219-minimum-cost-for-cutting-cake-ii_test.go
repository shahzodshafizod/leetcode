package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinimumCost$
func TestMinimumCost(t *testing.T) {
	for _, tc := range []struct {
		m             int
		n             int
		horizontalCut []int
		verticalCut   []int
		total         int64
	}{
		{m: 3, n: 2, horizontalCut: []int{1, 3}, verticalCut: []int{5}, total: 13},
		{m: 2, n: 2, horizontalCut: []int{7}, verticalCut: []int{4}, total: 15},
	} {
		total := minimumCost(tc.m, tc.n, tc.horizontalCut, tc.verticalCut)
		assert.Equal(t, tc.total, total)
	}
}
