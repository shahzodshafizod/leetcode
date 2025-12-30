package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestNumJewelsInStones$
func TestNumJewelsInStones(t *testing.T) {
	for _, tc := range []struct {
		jewels string
		stones string
		output int
	}{
		{"aA", "aAAbbbb", 3},           // Example 1: 'a', 'A', 'A' are jewels
		{"z", "ZZ", 0},                 // Example 2: case sensitive, no matches
		{"abc", "aabbccdd", 6},         // All a, b, c are jewels
		{"A", "AaAaAa", 3},             // Only uppercase A is jewel
		{"", "stones", 0},              // No jewels
		{"jewels", "", 0},              // No stones
		{"aAbBcC", "aAaAbBbBcCcC", 12}, // Multiple jewel types
		{"x", "xxxxXXXX", 4},           // Case sensitive 'x' only
	} {
		output := numJewelsInStones(tc.jewels, tc.stones)
		assert.Equal(t, tc.output, output)
	}
}
