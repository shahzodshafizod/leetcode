package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMincostToHireWorkers$
func TestMincostToHireWorkers(t *testing.T) {
	for _, tc := range []struct {
		quality []int
		wage    []int
		k       int
		money   float64
	}{
		{quality: []int{10, 20, 5}, wage: []int{70, 50, 30}, k: 2, money: 105.00000},
		{quality: []int{3, 1, 10, 10, 1}, wage: []int{4, 8, 2, 2, 7}, k: 3, money: 30.666666666666664},
	} {
		money := mincostToHireWorkers(tc.quality, tc.wage, tc.k)
		assert.InEpsilon(t, tc.money, money, 0.000000000000001)
	}
}
