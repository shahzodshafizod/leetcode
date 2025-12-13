package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxTotalReward$
func TestMaxTotalReward(t *testing.T) {
	for _, tc := range []struct {
		rewardValues []int
		output       int
	}{
		{rewardValues: []int{1, 1, 3, 3}, output: 4},     // Take 1, then 3: total = 4
		{rewardValues: []int{1, 6, 4, 3, 2}, output: 11}, // optimal is 1, 4, 6 = 11
		{rewardValues: []int{5}, output: 5},              // Take 5 (0 < 5): total = 5
		{rewardValues: []int{1}, output: 1},              // Take 1
		{rewardValues: []int{1, 2}, output: 3},           // Take 1, then 2: total = 3
		{rewardValues: []int{1, 2, 3}, output: 5},        // Take 2, 3: total = 5
	} {
		output := maxTotalReward(tc.rewardValues)
		assert.Equal(t, tc.output, output)
	}
}
