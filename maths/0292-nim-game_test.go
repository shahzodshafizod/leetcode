package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCanWinNim$
func TestCanWinNim(t *testing.T) {
	for _, tc := range []struct {
		n      int
		output bool
	}{
		{n: 1, output: true},
		{n: 2, output: true},
		{n: 3, output: true},
		{n: 4, output: false},
		{n: 5, output: true},
		{n: 6, output: true},
		{n: 7, output: true},
		{n: 8, output: false},
		{n: 12, output: false},
		{n: 15, output: true},
		{n: 100, output: false},
		{n: 101, output: true},
	} {
		output := canWinNim(tc.n)
		assert.Equal(t, tc.output, output)
	}
}
