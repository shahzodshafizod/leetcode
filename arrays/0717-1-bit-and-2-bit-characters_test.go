package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestIsOneBitCharacter$
func TestIsOneBitCharacter(t *testing.T) {
	for _, tc := range []struct {
		bits   []int
		output bool
	}{
		{bits: []int{1, 0, 0}, output: true},
		{bits: []int{1, 1, 1, 0}, output: false},
		{bits: []int{0}, output: true},
		{bits: []int{1, 0}, output: false},
		{bits: []int{0, 0}, output: true},
		{bits: []int{1, 1, 0}, output: true},
		{bits: []int{0, 1, 0}, output: false},
	} {
		output := isOneBitCharacter(tc.bits)
		assert.Equal(t, tc.output, output)
	}
}
