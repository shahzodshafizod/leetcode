package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestKInversePairs$
func TestKInversePairs(t *testing.T) {
	for _, tc := range []struct {
		n     int
		k     int
		count int
	}{
		{n: 3, k: 0, count: 1},
		{n: 3, k: 1, count: 2},
		{n: 3, k: 3, count: 1},
		{n: 4, k: 5, count: 3},
		{n: 4, k: 2, count: 5},
		// {n: 1000, k: 1000, count: 663677020},
	} {
		count := kInversePairs(tc.n, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
