package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestCountGoodIntegers$
func TestCountGoodIntegers(t *testing.T) {
	for _, tc := range []struct {
		n     int
		k     int
		count int64
	}{
		{n: 3, k: 5, count: 27},
		{n: 1, k: 4, count: 2},
		{n: 5, k: 6, count: 2468},
	} {
		count := countGoodIntegers(tc.n, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
