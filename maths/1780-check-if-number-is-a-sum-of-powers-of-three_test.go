package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCheckPowersOfThree$
func TestCheckPowersOfThree(t *testing.T) {
	for _, tc := range []struct {
		n        int
		possible bool
	}{
		{n: 12, possible: true},
		{n: 91, possible: true},
		{n: 21, possible: false},
	} {
		possible := checkPowersOfThree(tc.n)
		assert.Equal(t, tc.possible, possible)
	}
}
