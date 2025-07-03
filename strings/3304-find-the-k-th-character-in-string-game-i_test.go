package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestKthCharacter$
func TestKthCharacter(t *testing.T) {
	for _, tc := range []struct {
		k  int
		ch byte
	}{
		{k: 5, ch: 'b'},
		{k: 10, ch: 'c'},
	} {
		ch := kthCharacter(tc.k)
		assert.Equal(t, tc.ch, ch)
	}
}
