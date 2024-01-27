package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestNumOfMinutes$
func TestNumOfMinutes(t *testing.T) {
	for _, tc := range []struct {
		n          int
		headID     int
		manager    []int
		informTime []int
		minutes    int
	}{
		{n: 1, headID: 0, manager: []int{-1}, informTime: []int{0}, minutes: 0},
		{n: 6, headID: 2, manager: []int{2, 2, -1, 2, 2, 2}, informTime: []int{0, 0, 1, 0, 0, 0}, minutes: 1},
		{n: 7, headID: 6, manager: []int{1, 2, 3, 4, 5, 6, -1}, informTime: []int{0, 6, 5, 4, 3, 2, 1}, minutes: 21},
		{n: 8, headID: 4, manager: []int{2, 2, 4, 6, -1, 4, 4, 5}, informTime: []int{0, 0, 4, 0, 7, 3, 6, 0}, minutes: 13},
	} {
		minutes := numOfMinutes(tc.n, tc.headID, tc.manager, tc.informTime)
		assert.Equal(t, tc.minutes, minutes)
	}
}
