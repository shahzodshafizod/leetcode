package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCountGoodNumbers$
func TestCountGoodNumbers(t *testing.T) {
	for _, tc := range []struct {
		n     int64
		count int
	}{
		{n: 1, count: 5},
		{n: 4, count: 400},
		{n: 50, count: 564908303},
	} {
		count := countGoodNumbers(tc.n)
		assert.Equal(t, tc.count, count)
	}
}
