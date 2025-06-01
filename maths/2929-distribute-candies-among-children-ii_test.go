package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestDistributeCandies2929$
func TestDistributeCandies2929(t *testing.T) {
	for _, tc := range []struct {
		n     int
		limit int
		ways  int64
	}{
		{n: 5, limit: 2, ways: 3},
		{n: 3, limit: 3, ways: 10},
	} {
		ways := distributeCandies2929(tc.n, tc.limit)
		assert.Equal(t, tc.ways, ways)
	}
}
