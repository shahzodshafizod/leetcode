package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestMinValidSubstring$
func TestMinValidSubstring(t *testing.T) {
	for _, tc := range []struct {
		s      string
		p      string
		output int
	}{
		{s: "abaacbaecebce", p: "ba*c*ce", output: 8},
		{s: "baccbaadbc", p: "cc*baa*adb", output: -1},
		{s: "a", p: "**", output: 0},
		{s: "madlogic", p: "*adlogi*", output: 6},
		{s: "hy", p: "h**", output: 1},
	} {
		output := minValidSubstring(tc.s, tc.p)
		assert.Equal(t, tc.output, output)
	}
}
