package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxScore$
func TestMaxScore(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		score int
	}{
		{nums: []int{1, 2}, score: 1},
		{nums: []int{2, 3, 4, 6}, score: 8},
		{nums: []int{3, 4, 6, 8}, score: 11},
		{nums: []int{1, 2, 3, 4, 5, 6}, score: 14},
		{nums: []int{9, 17, 16, 15, 18, 13, 18, 20}, score: 91},
		{nums: []int{697035, 181412, 384958, 575458}, score: 869},
		{nums: []int{109497, 983516, 698308, 409009, 310455, 528595, 524079, 18036, 341150, 641864, 913962, 421869, 943382, 295019}, score: 527},
	} {
		score := maxScore(tc.nums)
		assert.Equal(t, tc.score, score)
	}
}
