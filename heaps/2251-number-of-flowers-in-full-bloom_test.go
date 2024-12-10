package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestFullBloomFlowers$
func TestFullBloomFlowers(t *testing.T) {
	for _, tc := range []struct {
		flowers [][]int
		people  []int
		answer  []int
	}{
		{flowers: [][]int{{1, 10}, {3, 3}}, people: []int{3, 3, 2}, answer: []int{2, 2, 1}},
		{flowers: [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}, people: []int{2, 3, 7, 11}, answer: []int{1, 2, 2, 2}},
	} {
		answer := fullBloomFlowers(tc.flowers, tc.people)
		assert.Equal(t, tc.answer, answer)
	}
}
