package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCountSubstrings$
func TestCountSubstrings(t *testing.T) {
	for _, tc := range []struct {
		s     string
		count int
	}{
		{s: "abc", count: 3},
		{s: "aaa", count: 6},
		{s: "aaaa", count: 10},
		{s: "a", count: 1},
	} {
		count := countSubstrings(tc.s)
		assert.Equal(t, tc.count, count)
	}
}
