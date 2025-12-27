package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestUniqueLetterString$
func TestUniqueLetterString(t *testing.T) {
	for _, tc := range []struct {
		s      string
		output int
	}{
		{s: "ABC", output: 10},
		{s: "ABA", output: 8},
		{s: "LEETCODE", output: 92},
		{s: "A", output: 1},
		{s: "AA", output: 2},
	} {
		output := uniqueLetterString(tc.s)
		assert.Equal(t, tc.output, output)
	}
}
