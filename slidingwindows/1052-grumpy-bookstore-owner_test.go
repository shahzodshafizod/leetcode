package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestMaxSatisfied$
func TestMaxSatisfied(t *testing.T) {
	for _, tc := range []struct {
		customers []int
		grumpy    []int
		minutes   int
		satisfied int
	}{
		{customers: []int{1, 0, 1, 2, 1, 1, 7, 5}, grumpy: []int{0, 1, 0, 1, 0, 1, 0, 1}, minutes: 3, satisfied: 16},
		{customers: []int{1}, grumpy: []int{0}, minutes: 1, satisfied: 1},
		{customers: []int{10, 1, 7}, grumpy: []int{0, 0, 0}, minutes: 2, satisfied: 18},
		{customers: []int{1, 2, 1, 5, 3, 2, 1}, grumpy: []int{1, 0, 1, 0, 1, 0, 1}, minutes: 2, satisfied: 12},
		{customers: []int{4, 10, 10}, grumpy: []int{1, 1, 0}, minutes: 2, satisfied: 24},
	} {
		satisfied := maxSatisfied(tc.customers, tc.grumpy, tc.minutes)
		assert.Equal(t, tc.satisfied, satisfied)
	}
}
