package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaximumTripletValue$
func TestMaximumTripletValue(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		value int64
	}{
		{nums: []int{12, 6, 1, 2, 7}, value: 77},
		{nums: []int{1, 10, 3, 4, 19}, value: 133},
		{nums: []int{1, 2, 3}, value: 0},
	} {
		value := maximumTripletValue(tc.nums)
		assert.Equal(t, tc.value, value)
	}
}
