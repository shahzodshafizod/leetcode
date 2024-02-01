package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestKnightProbability$
func TestKnightProbability(t *testing.T) {
	for _, tc := range []struct {
		n           int
		k           int
		row         int
		column      int
		probability float64
	}{
		{n: 3, k: 2, row: 0, column: 0, probability: 0.06250},
		{n: 1, k: 0, row: 0, column: 0, probability: 1.00000},
		{n: 6, k: 2, row: 2, column: 2, probability: 0.53125},
		{n: 6, k: 3, row: 2, column: 2, probability: 0.359375},
		{n: 2, k: 3, row: 1, column: 1, probability: 0.00000},
		{n: 2, k: 0, row: 0, column: 1, probability: 1.00000},
	} {
		probability := knightProbability(tc.n, tc.k, tc.row, tc.column)
		assert.Equal(t, tc.probability, probability)
	}
}
