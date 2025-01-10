package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestWordSubsets$
func TestWordSubsets(t *testing.T) {
	for _, tc := range []struct {
		words1  []string
		words2  []string
		subsets []string
	}{
		{words1: []string{"amazon", "apple", "facebook", "google", "leetcode"}, words2: []string{"e", "o"}, subsets: []string{"facebook", "google", "leetcode"}},
		{words1: []string{"amazon", "apple", "facebook", "google", "leetcode"}, words2: []string{"l", "e"}, subsets: []string{"apple", "google", "leetcode"}},
	} {
		subsets := wordSubsets(tc.words1, tc.words2)
		assert.Equal(t, tc.subsets, subsets)
	}
}
