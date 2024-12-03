package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestLongestObstacleCourseAtEachPosition$
func TestLongestObstacleCourseAtEachPosition(t *testing.T) {
	for _, tc := range []struct {
		obstacles []int
		ans       []int
	}{
		{obstacles: []int{1, 2, 3, 2}, ans: []int{1, 2, 3, 3}},
		{obstacles: []int{2, 2, 1}, ans: []int{1, 2, 1}},
		{obstacles: []int{3, 1, 5, 6, 4, 2}, ans: []int{1, 1, 2, 3, 2, 2}},
		{obstacles: []int{5, 1, 5, 5, 1, 3, 4, 5, 1, 4}, ans: []int{1, 1, 2, 3, 2, 3, 4, 5, 3, 5}},
	} {
		ans := longestObstacleCourseAtEachPosition(tc.obstacles)
		assert.Equal(t, tc.ans, ans)
	}
}
