package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestPutMarbles$
func TestPutMarbles(t *testing.T) {
	for _, tc := range []struct {
		weights    []int
		k          int
		difference int64
	}{
		{weights: []int{1, 3, 5, 1}, k: 2, difference: 4},
		{weights: []int{1, 3}, k: 2, difference: 0},
		{weights: []int{54, 6, 34, 66, 63, 52, 39, 62, 46, 75, 28, 65, 18, 37, 18, 13, 33, 69, 19, 40, 13, 10, 43, 61, 72}, k: 4, difference: 289},
		{weights: []int{1, 4, 2, 5, 2}, k: 3, difference: 3},
		{weights: []int{1, 9, 8, 7, 1, 1, 1}, k: 3, difference: 28},
	} {
		difference := putMarbles(tc.weights, tc.k)
		assert.Equal(t, tc.difference, difference)
	}
}
