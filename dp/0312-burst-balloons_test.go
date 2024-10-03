package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxCoins$
func TestMaxCoins(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		coins int
	}{
		{nums: []int{3, 1, 5, 8}, coins: 167},
		{nums: []int{1, 5}, coins: 10},
		{nums: []int{9, 76, 64, 21, 97, 60}, coins: 1086136},
	} {
		coins := maxCoins(tc.nums)
		assert.Equal(t, tc.coins, coins)
	}
}
