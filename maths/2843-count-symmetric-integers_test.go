package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCountSymmetricIntegers$
func TestCountSymmetricIntegers(t *testing.T) {
	for _, tc := range []struct {
		low   int
		high  int
		count int
	}{
		{low: 1, high: 100, count: 9},
		{low: 1200, high: 1230, count: 4},
	} {
		count := countSymmetricIntegers(tc.low, tc.high)
		assert.Equal(t, tc.count, count)
	}
}
