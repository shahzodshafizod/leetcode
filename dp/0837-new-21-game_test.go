package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNew21Game$
func TestNew21Game(t *testing.T) {
	for _, tc := range []struct {
		n int
		k int
		maxPts int
		probability float64
	}{
		{n: 10, k: 1, maxPts: 10, probability: 1.00000},
		{n: 6, k: 1, maxPts: 10, probability: 0.60000},
		{n: 21, k: 17, maxPts: 10, probability: 0.73278},
	} {
		probability := new21Game(tc.n, tc.k, tc.maxPts)
		assert.InDelta(t, tc.probability, probability, 0.00001)
	}
}
