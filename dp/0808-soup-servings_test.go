package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestSoupServings$
func TestSoupServings(t *testing.T) {
	for _, tc := range []struct {
		n           int
		probability float64
	}{
		{n: 50, probability: 0.62500},
		{n: 100, probability: 0.71875},
	} {
		probability := soupServings(tc.n)
		assert.InDelta(t, tc.probability, probability, 0.00001)
	}
}
