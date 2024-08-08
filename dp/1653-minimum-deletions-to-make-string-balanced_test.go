package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinimumDeletions$
func TestMinimumDeletions(t *testing.T) {
	for _, tc := range []struct {
		s         string
		deletions int
	}{
		{s: "aababbab", deletions: 2},
		{s: "bbaaaaabb", deletions: 2},
		{s: "a", deletions: 0},
		{s: "b", deletions: 0},
		{s: "ababaaaabbbbbaaababbbbbbaaabbaababbabbbbaabbbbaabbabbabaabbbababaa", deletions: 25},
		{s: "baabab", deletions: 2},
	} {
		deletions := minimumDeletions(tc.s)
		assert.Equal(t, tc.deletions, deletions)
	}
}
