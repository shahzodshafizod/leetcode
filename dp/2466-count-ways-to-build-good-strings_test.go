package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCountGoodStrings$
func TestCountGoodStrings(t *testing.T) {
	for _, tc := range []struct {
		low   int
		high  int
		zero  int
		one   int
		count int
	}{
		{low: 3, high: 3, zero: 1, one: 1, count: 8},
		{low: 2, high: 3, zero: 1, one: 2, count: 5},
		{low: 1, high: 1, zero: 1, one: 1, count: 2},
		{low: 100000, high: 100000, zero: 100000, one: 100000, count: 2},
		// {low: 200, high: 200, zero: 10, one: 1, count: 764262396},
		// {low: 1, high: 100000, zero: 1, one: 1, count: 215447031},
		// {low: 200, high: 200, zero: 10, one: 1, count: 764262396},
		// {low: 277, high: 99991, zero: 31, one: 19, count: 271302617},
		// {low: 2657, high: 9601, zero: 1, one: 2, count: 267484946},
		// {low: 2, high: 8699, zero: 1, one: 1, count: 987490208},
	} {
		count := countGoodStrings(tc.low, tc.high, tc.zero, tc.one)
		assert.Equal(t, tc.count, count)
	}
}
