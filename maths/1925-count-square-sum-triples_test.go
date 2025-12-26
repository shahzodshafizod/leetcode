package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCountTriples$
func TestCountTriples(t *testing.T) {
	for _, tc := range []struct {
		n     int
		count int
	}{
		{n: 5, count: 2},  // (3,4,5), (4,3,5)
		{n: 10, count: 4}, // (3,4,5), (4,3,5), (6,8,10), (8,6,10)
		{n: 1, count: 0},
		{n: 2, count: 0},
		{n: 15, count: 8}, // 8 triples
	} {
		count := countTriples(tc.n)
		assert.Equal(t, tc.count, count)
	}
}
