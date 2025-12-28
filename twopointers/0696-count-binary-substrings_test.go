package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestCountBinarySubstrings$
func TestCountBinarySubstrings(t *testing.T) {
	for _, tc := range []struct {
		s      string
		result int
	}{
		{s: "00110011", result: 6},
		{s: "10101", result: 4},
		{s: "00110", result: 3},
		{s: "000111", result: 3},
		{s: "01", result: 1},
		{s: "0011", result: 2},
		{s: "00100", result: 2},
		{s: "001101", result: 4},
	} {
		result := countBinarySubstrings(tc.s)
		assert.Equal(t, tc.result, result)
	}
}
