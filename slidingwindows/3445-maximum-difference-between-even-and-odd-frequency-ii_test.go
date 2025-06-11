package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestMaxDifference$
func TestMaxDifference(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		diff int
	}{
		{s: "12233", k: 4, diff: -1},
		{s: "1122211", k: 3, diff: 1},
		{s: "110", k: 3, diff: -1},
		{s: "2222130", k: 2, diff: -1},
		{s: "11131340", k: 8, diff: -1},
		{s: "111112222233333", k: 5, diff: 3},
	} {
		diff := maxDifference(tc.s, tc.k)
		assert.Equal(t, tc.diff, diff)
	}
}
