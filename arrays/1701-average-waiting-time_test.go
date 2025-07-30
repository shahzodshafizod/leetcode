package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestAverageWaitingTime$
func TestAverageWaitingTime(t *testing.T) {
	for _, tc := range []struct {
		customers       [][]int
		averageWaitTime float64
	}{
		{customers: [][]int{{1, 2}, {2, 5}, {4, 3}}, averageWaitTime: 5.00000},
		{customers: [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}, averageWaitTime: 3.25000},
	} {
		averageWaitTime := averageWaitingTime(tc.customers)
		assert.InEpsilon(t, tc.averageWaitTime, averageWaitTime, 0.00001)
	}
}
