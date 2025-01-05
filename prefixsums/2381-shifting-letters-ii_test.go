package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestShiftingLetters$
func TestShiftingLetters(t *testing.T) {
	for _, tc := range []struct {
		s       string
		shifts  [][]int
		shifted string
	}{
		{s: "abc", shifts: [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}, shifted: "ace"},
		{s: "dztz", shifts: [][]int{{0, 0, 0}, {1, 1, 1}}, shifted: "catz"},
	} {
		shifted := shiftingLetters(tc.s, tc.shifts)
		assert.Equal(t, tc.shifted, shifted)
	}
}
