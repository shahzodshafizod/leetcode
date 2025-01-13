package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestCanBeValid$
func TestCanBeValid(t *testing.T) {
	for _, tc := range []struct {
		s      string
		locked string
		can    bool
	}{
		{s: "))()))", locked: "010100", can: true},
		{s: "()()", locked: "0000", can: true},
		{s: ")", locked: "0", can: false},
		{s: "())()))()(()(((())(()()))))((((()())(())", locked: "1011101100010001001011000000110010100101", can: true},
	} {
		can := canBeValid(tc.s, tc.locked)
		assert.Equal(t, tc.can, can)
	}
}
