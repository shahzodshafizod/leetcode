package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestRomanToInt$
func TestRomanToInt(t *testing.T) {
	for _, tc := range []struct {
		s   string
		num int
	}{
		{s: "III", num: 3},
		{s: "LVIII", num: 58},
		{s: "MCMXCIV", num: 1994},
		{s: "II", num: 2},
		{s: "XII", num: 12},
		{s: "XXVII", num: 27},
		{s: "IV", num: 4},
		{s: "IX", num: 9},
	} {
		num := romanToInt(tc.s)
		assert.Equal(t, tc.num, num)
	}
}
