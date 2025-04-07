package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCanPartition$
func TestCanPartition(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		can  bool
	}{
		{nums: []int{1, 5, 11, 5}, can: true},
		{nums: []int{1, 2, 3, 5}, can: false},
	} {
		can := canPartition(tc.nums)
		assert.Equal(t, tc.can, can)
	}
}
