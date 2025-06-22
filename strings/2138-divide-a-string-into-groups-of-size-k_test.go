package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestDivideString$
func TestDivideString(t *testing.T) {
	for _, tc := range []struct {
		s          string
		k          int
		fill       byte
		partitions []string
	}{
		{s: "abcdefghi", k: 3, fill: 'x', partitions: []string{"abc", "def", "ghi"}},
		{s: "abcdefghij", k: 3, fill: 'x', partitions: []string{"abc", "def", "ghi", "jxx"}},
	} {
		partitions := divideString(tc.s, tc.k, tc.fill)
		assert.Equal(t, tc.partitions, partitions)
	}
}
