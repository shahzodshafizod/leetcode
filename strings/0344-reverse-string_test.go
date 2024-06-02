package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestReverseString$
func TestReverseString(t *testing.T) {
	for _, tc := range []struct {
		s         []byte
		reversedS []byte
	}{
		{s: []byte{'h', 'e', 'l', 'l', 'o'}, reversedS: []byte{'o', 'l', 'l', 'e', 'h'}},
		{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}, reversedS: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	} {
		reverseString(tc.s)
		assert.Equal(t, tc.reversedS, tc.s)
	}
}
