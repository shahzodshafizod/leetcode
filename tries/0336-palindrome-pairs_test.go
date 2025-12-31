package tries

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestPalindromePairs$
func TestPalindromePairs(t *testing.T) {
	for _, tc := range []struct {
		words  []string
		output [][]int
	}{
		{
			words:  []string{"abcd", "dcba", "lls", "s", "sssll"},
			output: [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}},
		},
		{
			words:  []string{"bat", "tab", "cat"},
			output: [][]int{{0, 1}, {1, 0}},
		},
		{
			words:  []string{"a", ""},
			output: [][]int{{0, 1}, {1, 0}},
		},
		{
			words:  []string{"a", "abc", "aba", ""},
			output: [][]int{{0, 3}, {2, 3}, {3, 0}, {3, 2}},
		},
	} {
		output := palindromePairs(tc.words)
		// Sort for comparison since order may vary
		sort.Slice(output, func(i, j int) bool {
			if output[i][0] != output[j][0] {
				return output[i][0] < output[j][0]
			}

			return output[i][1] < output[j][1]
		})
		sort.Slice(tc.output, func(i, j int) bool {
			if tc.output[i][0] != tc.output[j][0] {
				return tc.output[i][0] < tc.output[j][0]
			}

			return tc.output[i][1] < tc.output[j][1]
		})
		assert.Equal(t, tc.output, output)
	}
}
