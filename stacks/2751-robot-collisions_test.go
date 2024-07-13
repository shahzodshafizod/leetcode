package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestSurvivedRobotsHealths$
func TestSurvivedRobotsHealths(t *testing.T) {
	for _, tc := range []struct {
		positions  []int
		healths    []int
		directions string
		survivors  []int
	}{
		{positions: []int{5, 4, 3, 2, 1}, healths: []int{2, 17, 9, 15, 10}, directions: "RRRRR", survivors: []int{2, 17, 9, 15, 10}},
		{positions: []int{3, 5, 2, 6}, healths: []int{10, 10, 15, 12}, directions: "RLRL", survivors: []int{14}},
		{positions: []int{1, 2, 5, 6}, healths: []int{10, 10, 11, 11}, directions: "RLRL", survivors: []int{}},
		{positions: []int{3, 47}, healths: []int{46, 26}, directions: "LR", survivors: []int{46, 26}},
		{positions: []int{37, 35}, healths: []int{16, 19}, directions: "RL", survivors: []int{16, 19}},
		{positions: []int{13, 3}, healths: []int{17, 2}, directions: "LR", survivors: []int{16}},
		{positions: []int{3, 40}, healths: []int{49, 11}, directions: "LL", survivors: []int{49, 11}},
	} {
		survivors := survivedRobotsHealths(tc.positions, tc.healths, tc.directions)
		assert.Equal(t, tc.survivors, survivors)
	}
}
