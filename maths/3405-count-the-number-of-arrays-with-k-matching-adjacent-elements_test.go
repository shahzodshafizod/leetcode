package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCountGoodArrays$
func TestCountGoodArrays(t *testing.T) {
	for _, tc := range []struct {
		n     int
		m     int
		k     int
		count int
	}{
		{n: 3, m: 2, k: 1, count: 4},
		{n: 4, m: 2, k: 2, count: 6},
		{n: 5, m: 2, k: 0, count: 2},
		{n: 5581, m: 58624, k: 4766, count: 846088010},
	} {
		count := countGoodArrays(tc.n, tc.m, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
