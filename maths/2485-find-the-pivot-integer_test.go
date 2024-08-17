package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestPivotInteger$
func TestPivotInteger(t *testing.T) {
	for _, tc := range []struct {
		n     int
		pivot int
	}{
		{n: 8, pivot: 6},
		{n: 1, pivot: 1},
		{n: 4, pivot: -1},
	} {
		pivot := pivotInteger(tc.n)
		assert.Equal(t, tc.pivot, pivot)
	}
}
