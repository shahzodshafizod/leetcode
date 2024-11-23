package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestJobScheduling$
func TestJobScheduling(t *testing.T) {
	for _, tc := range []struct {
		startTime []int
		endTime   []int
		profit    []int
		maxProfit int
	}{
		{startTime: []int{7779}, endTime: []int{13973}, profit: []int{4658}, maxProfit: 4658},
		{startTime: []int{1, 1, 1}, endTime: []int{2, 3, 4}, profit: []int{5, 6, 4}, maxProfit: 6},
		{startTime: []int{3559, 5372}, endTime: []int{8592, 11236}, profit: []int{8968, 2073}, maxProfit: 8968},
		{startTime: []int{1, 2, 3, 3}, endTime: []int{3, 4, 5, 6}, profit: []int{50, 10, 40, 70}, maxProfit: 120},
		{startTime: []int{1, 2, 3, 4, 6}, endTime: []int{3, 5, 10, 6, 9}, profit: []int{20, 20, 100, 70, 60}, maxProfit: 150},
		{startTime: []int{497, 9813, 5338, 2256, 9062}, endTime: []int{7676, 18960, 14385, 9924, 12218}, profit: []int{2712, 9026, 603, 5918, 9964}, maxProfit: 12676},
		{startTime: []int{6, 24, 45, 27, 13, 43, 47, 36, 14, 11, 11, 12}, endTime: []int{31, 27, 48, 46, 44, 46, 50, 49, 24, 42, 13, 27}, profit: []int{14, 4, 16, 12, 20, 3, 18, 6, 9, 1, 2, 8}, maxProfit: 45},
		{startTime: []int{8882, 7030, 7842, 8505, 8586, 2257, 8970, 1669, 1124, 3159}, endTime: []int{18708, 10088, 7955, 9997, 9173, 4828, 13306, 5053, 1380, 8825}, profit: []int{9933, 967, 5417, 1835, 2746, 3901, 773, 2248, 5062, 3712}, maxProfit: 24313},
		// {startTime: []int{96, 815, 776, 701, 4773, 3771, 6725, 8841, 4289, 4992, 9852, 6641, 6065, 3366, 6415, 7684, 9429, 7620, 8880, 6509}, endTime: []int{3816, 7818, 8658, 3359, 6845, 12795, 12943, 13524, 8099, 9911, 16365, 7667, 10043, 9038, 7822, 14342, 13581, 12347, 10690, 11189}, profit: []int{8628, 8418, 9036, 7717, 1828, 1498, 5657, 1156, 4749, 2528, 8675, 4130, 3791, 6585, 655, 3686, 8462, 9630, 2263, 9595}, maxProfit: 22977},
	} {
		maxProfit := jobScheduling(tc.startTime, tc.endTime, tc.profit)
		assert.Equal(t, tc.maxProfit, maxProfit)
	}
}
