package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestMaxPriorityQueueSort$
func TestMaxPriorityQueueSort(t *testing.T) {
	for _, tc := range []struct {
		array  []int
		sorted []int
	}{
		{
			array:  []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5},
			sorted: []int{45, 11, 7, 7, 6, 6, 6, 6, 5, 5, 5, 4, 4, 4, 4, 3, 2, 2, 1, 1, 0},
		},
	} {
		var maxHeap = NewPQSort(
			len(tc.array),
			func(i, j int) bool { return tc.array[i] > tc.array[j] },
			func(i, j int) { tc.array[i], tc.array[j] = tc.array[j], tc.array[i] },
		)
		maxHeap.Sort()
		assert.Equal(t, tc.sorted, tc.array)
	}
}

// go test -v -count=1 ./design/ -run ^TestMinPriorityQueueSort$
func TestMinPriorityQueueSort(t *testing.T) {
	for _, tc := range []struct {
		array  []int
		sorted []int
	}{
		{
			array:  []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5},
			sorted: []int{0, 1, 1, 2, 2, 3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 6, 7, 7, 11, 45},
		},
	} {
		var minHeap = NewPQSort(
			len(tc.array),
			func(i, j int) bool { return tc.array[i] < tc.array[j] },
			func(i int, j int) { tc.array[i], tc.array[j] = tc.array[j], tc.array[i] },
		)
		minHeap.Sort()
		assert.Equal(t, tc.sorted, tc.array)
	}
}
