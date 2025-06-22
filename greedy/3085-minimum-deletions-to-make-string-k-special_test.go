package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinimumDeletions$
func TestMinimumDeletions(t *testing.T) {
	for _, tc := range []struct {
		word      string
		k         int
		deletions int
	}{
		{word: "aabcaba", k: 0, deletions: 3},
		{word: "dabdcbdcdcd", k: 2, deletions: 2},
		{word: "aaabaaa", k: 2, deletions: 1},
	} {
		deletions := minimumDeletions(tc.word, tc.k)
		assert.Equal(t, tc.deletions, deletions)
	}
}
