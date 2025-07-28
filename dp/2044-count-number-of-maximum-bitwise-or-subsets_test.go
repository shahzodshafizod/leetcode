package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCountMaxOrSubsets$
func TestCountMaxOrSubsets(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int
	}{
		{nums: []int{3, 1}, count: 2},
		{nums: []int{2, 2, 2}, count: 7},
		{nums: []int{3, 2, 1, 5}, count: 6},
		{nums: []int{12007}, count: 1},
		{nums: []int{39317, 28735, 71230, 59313, 19954}, count: 5},
		{nums: []int{35569, 91997, 54930, 66672, 12363}, count: 6},
		// {nums: []int{32, 39, 37, 31, 42, 38, 29, 43, 40, 36, 44, 35, 41, 33, 34, 30}, count: 57083},
		// {nums: []int{63609, 94085, 69432, 15248, 22060, 76843, 84075, 835, 23463, 66399, 95031, 22676, 92115}, count: 3772},
		// {nums: []int{6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890}, count: 65535},
		// {nums: []int{89260, 58129, 16949, 64128, 9782, 26664, 96968, 59838, 27627, 68561, 4026, 91345, 26966, 28876, 46276, 19878}, count: 52940},
		// {nums: []int{90193, 56697, 77229, 72927, 74728, 93652, 70751, 32415, 94774, 9067, 14758, 59835, 26047, 36393, 60530, 64649}, count: 60072},
	} {
		count := countMaxOrSubsets(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
