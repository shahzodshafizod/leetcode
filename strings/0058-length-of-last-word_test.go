package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestLengthOfLastWord$
func TestLengthOfLastWord(t *testing.T) {
	for _, tc := range []struct {
		s      string
		length int
	}{
		{s: "Hello World", length: 5},
		{s: "   fly me   to   the moon  ", length: 4},
		{s: "luffy is still joyboy", length: 6},
		{s: "a ", length: 1},
		{s: "a", length: 1},
	} {
		length := lengthOfLastWord(tc.s)
		assert.Equal(t, tc.length, length)
	}
}
