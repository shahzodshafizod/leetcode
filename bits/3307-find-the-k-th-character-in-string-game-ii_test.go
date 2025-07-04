package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestKthCharacter$
func TestKthCharacter(t *testing.T) {
	for _, tc := range []struct {
		k          int64
		operations []int
		char       byte
	}{
		{k: 5, operations: []int{0, 0, 0}, char: 'a'},
		{k: 10, operations: []int{0, 1, 0, 1}, char: 'b'},
		{k: 33031255, operations: []int{0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1}, char: 'j'},
		{k: 14493383, operations: []int{0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1}, char: 'i'},
		{k: 32961520, operations: []int{1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1}, char: 'k'},
	} {
		char := kthCharacter(tc.k, tc.operations)
		assert.Equal(t, tc.char, char)
	}
}
