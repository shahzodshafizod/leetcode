package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestColorTheGrid$
func TestColorTheGrid(t *testing.T) {
	for _, tc := range []struct {
		m     int
		n     int
		count int
	}{
		{m: 1, n: 1, count: 3},
		{m: 1, n: 2, count: 6},
		{m: 5, n: 5, count: 580986},
		// {m: 5, n: 1000, count: 408208448},
	} {
		count := colorTheGrid(tc.m, tc.n)
		assert.Equal(t, tc.count, count)
	}
}
