package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinNumberOperations$
func TestMinNumberOperations(t *testing.T) {
	for _, tc := range []struct {
		target []int
		ops    int
	}{
		{target: []int{1, 2, 3, 2, 1}, ops: 3},
		{target: []int{3, 1, 1, 2}, ops: 4},
		{target: []int{3, 1, 5, 4, 2}, ops: 7},
	} {
		ops := minNumberOperations(tc.target)
		assert.Equal(t, tc.ops, ops)
	}
}
