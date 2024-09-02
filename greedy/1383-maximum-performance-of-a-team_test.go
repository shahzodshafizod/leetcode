package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaxPerformance$
func TestMaxPerformance(t *testing.T) {
	for _, tc := range []struct {
		n           int
		speed       []int
		efficiency  []int
		k           int
		performance int
	}{
		{n: 6, speed: []int{2, 10, 3, 1, 5, 8}, efficiency: []int{5, 4, 3, 9, 7, 2}, k: 2, performance: 60},
		{n: 6, speed: []int{2, 10, 3, 1, 5, 8}, efficiency: []int{5, 4, 3, 9, 7, 2}, k: 3, performance: 68},
		{n: 6, speed: []int{2, 10, 3, 1, 5, 8}, efficiency: []int{5, 4, 3, 9, 7, 2}, k: 4, performance: 72},
		{n: 3, speed: []int{2, 8, 2}, efficiency: []int{2, 7, 1}, k: 2, performance: 56},
	} {
		performance := maxPerformance(tc.n, tc.speed, tc.efficiency, tc.k)
		assert.Equal(t, tc.performance, performance)
	}
}
