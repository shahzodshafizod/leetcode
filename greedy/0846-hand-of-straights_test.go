package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestIsNStraightHand$
func TestIsNStraightHand(t *testing.T) {
	for _, tc := range []struct {
		hand      []int
		groupSize int
		is        bool
	}{
		{hand: []int{1, 2, 3, 6, 2, 3, 4, 7, 8}, groupSize: 3, is: true},
		{hand: []int{1, 2, 3, 4, 5}, groupSize: 4, is: false},
		{hand: []int{2, 1}, groupSize: 2, is: true},
		{hand: []int{8, 8, 9, 7, 7, 7, 6, 7, 10, 6}, groupSize: 2, is: true},
		{hand: []int{8, 10, 12}, groupSize: 3, is: false},
		{hand: []int{1, 9, 17, 25, 33, 41, 49}, groupSize: 7, is: false},
	} {
		is := isNStraightHand(tc.hand, tc.groupSize)
		assert.Equal(t, tc.is, is)
	}
}
