package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinDeletionSize$
func TestMinDeletionSize(t *testing.T) {
	for _, tc := range []struct {
		strs   []string
		output int
	}{
		{[]string{"babca", "bbazb"}, 3},
		{[]string{"edcba"}, 4},
		{[]string{"ghi", "def", "abc"}, 0},
		{[]string{"a", "b"}, 0},
		{[]string{"zyx", "wvu", "tsr"}, 2},
		{[]string{"abc", "bcd", "cde"}, 0},
		{[]string{"ca", "bb", "ac"}, 1},
	} {
		output := minDeletionSize(tc.strs)
		assert.Equal(t, tc.output, output)
	}
}
