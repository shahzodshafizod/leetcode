package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestBuildArray$
func TestBuildArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		ans  []int
	}{
		{nums: []int{0, 2, 1, 5, 3, 4}, ans: []int{0, 1, 2, 4, 5, 3}},
		{nums: []int{5, 0, 1, 2, 3, 4}, ans: []int{4, 5, 0, 1, 2, 3}},
	} {
		ans := buildArray(tc.nums)
		assert.Equal(t, tc.ans, ans)
	}
}
