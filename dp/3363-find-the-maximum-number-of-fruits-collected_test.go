package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxCollectedFruits$
func TestMaxCollectedFruits(t *testing.T) {
	for _, tc := range []struct {
		fruits [][]int
		count  int
	}{
		{fruits: [][]int{{1, 2, 3, 4}, {5, 6, 8, 7}, {9, 10, 11, 12}, {13, 14, 15, 16}}, count: 100},
		{fruits: [][]int{{1, 1}, {1, 1}}, count: 4},
	} {
		count := maxCollectedFruits(tc.fruits)
		assert.Equal(t, tc.count, count)
	}
}
