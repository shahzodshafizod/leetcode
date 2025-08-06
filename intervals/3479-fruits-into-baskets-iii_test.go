package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./intervals/ -run ^TestNumOfUnplacedFruits$
func TestNumOfUnplacedFruits(t *testing.T) {
	for _, tc := range []struct {
		fruits  []int
		baskets []int
		count   int
	}{
		{fruits: []int{4, 2, 5}, baskets: []int{3, 5, 4}, count: 1},
		{fruits: []int{3, 6, 1}, baskets: []int{6, 4, 7}, count: 0},
	} {
		count := numOfUnplacedFruits(tc.fruits, tc.baskets)
		assert.Equal(t, tc.count, count)
	}
}
