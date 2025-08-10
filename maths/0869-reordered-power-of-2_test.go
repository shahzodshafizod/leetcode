package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestReorderedPowerOf2$
func TestReorderedPowerOf2(t *testing.T) {
	for _, tc := range []struct {
		n int
		can bool
	}{
		{n: 1, can: true},
		{n: 10, can: false},
		{n: 24, can: false},
	} {
		can := reorderedPowerOf2(tc.n)
		assert.Equal(t, tc.can, can)
	}
}
