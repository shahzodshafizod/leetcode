package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestNumWaterBottles$
func TestNumWaterBottles(t *testing.T) {
	for _, tc := range []struct {
		numBottles  int
		numExchange int
		consumed    int
	}{
		{numBottles: 9, numExchange: 3, consumed: 13},
		{numBottles: 15, numExchange: 4, consumed: 19},
	} {
		consumed := numWaterBottles(tc.numBottles, tc.numExchange)
		assert.Equal(t, tc.consumed, consumed)
	}
}
