package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaximumHappinessSum$
func TestMaximumHappinessSum(t *testing.T) {
	for _, tc := range []struct {
		happiness []int
		k         int
		sum       int64
	}{
		{happiness: []int{1, 2, 3}, k: 2, sum: 4},
		{happiness: []int{1, 1, 1, 1}, k: 2, sum: 1},
		{happiness: []int{2, 3, 4, 5}, k: 1, sum: 5},
		{happiness: []int{12, 1, 42}, k: 3, sum: 53},
	} {
		sum := maximumHappinessSum(tc.happiness, tc.k)
		assert.Equal(t, tc.sum, sum)
	}
}
