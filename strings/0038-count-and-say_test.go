package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestCountAndSay$
func TestCountAndSay(t *testing.T) {
	for _, tc := range []struct {
		n   int
		rle string
	}{
		{n: 4, rle: "1211"},
		{n: 1, rle: "1"},
	} {
		rle := countAndSay(tc.n)
		assert.Equal(t, tc.rle, rle)
	}
}
