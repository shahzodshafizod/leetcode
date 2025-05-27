package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestDifferenceOfSums$
func TestDifferenceOfSums(t *testing.T) {
	for _, tc := range []struct {
		n    int
		m    int
		diff int
	}{
		{n: 10, m: 3, diff: 19},
		{n: 5, m: 6, diff: 15},
		{n: 5, m: 1, diff: -15},
	} {
		diff := differenceOfSums(tc.n, tc.m)
		assert.Equal(t, tc.diff, diff)
	}
}
