package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMinDays$
func TestMinDays(t *testing.T) {
	for _, tc := range []struct {
		bloomDay []int
		m        int
		k        int
		days     int
	}{
		{bloomDay: []int{1, 10, 3, 10, 2}, m: 3, k: 1, days: 3},
		{bloomDay: []int{1, 10, 3, 10, 2}, m: 3, k: 2, days: -1},
		{bloomDay: []int{7, 7, 7, 7, 12, 7, 7}, m: 2, k: 3, days: 12},
		{bloomDay: []int{3, 2, 6, 1, 1, 2, 4}, m: 2, k: 2, days: 3},
	} {
		days := minDays(tc.bloomDay, tc.m, tc.k)
		assert.Equal(t, tc.days, days)
	}
}
