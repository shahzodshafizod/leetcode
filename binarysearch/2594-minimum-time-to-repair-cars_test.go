package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestRepairCars$
func TestRepairCars(t *testing.T) {
	for _, tc := range []struct {
		ranks []int
		cars  int
		time  int64
	}{
		{ranks: []int{4, 2, 3, 1}, cars: 10, time: 16},
		{ranks: []int{5, 1, 8}, cars: 6, time: 16},
	} {
		time := repairCars(tc.ranks, tc.cars)
		assert.Equal(t, tc.time, time)
	}
}
