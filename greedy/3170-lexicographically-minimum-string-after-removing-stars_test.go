package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestClearStars$
func TestClearStars(t *testing.T) {
	for _, tc := range []struct {
		s       string
		cleared string
	}{
		{s: "aaba*", cleared: "aab"},
		{s: "abc", cleared: "abc"},
		{s: "abc*de*fgh*", cleared: "defgh"},
		{s: "a*b*c*d*", cleared: ""},
		{s: "abcde*f*", cleared: "cdef"},
		{s: "abc*", cleared: "bc"},
		{s: "aaaa***", cleared: "a"},
		{s: "de*", cleared: "e"},
	} {
		cleared := clearStars(tc.s)
		assert.Equal(t, tc.cleared, cleared)
	}
}
