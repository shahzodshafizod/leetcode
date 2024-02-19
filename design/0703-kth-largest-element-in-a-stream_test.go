package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestKthLargest$
func TestKthLargest(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]any
		output   []any
	}{
		{
			commands: []string{"KthLargest", "add", "add", "add", "add", "add"},
			values:   [][]any{{3, []any{4, 5, 8, 2}}, {3}, {5}, {10}, {9}, {4}},
			output:   []any{nil, 4, 5, 5, 8, 8},
		},
		{
			commands: []string{"KthLargest", "add", "add", "add", "add", "add"},
			values:   [][]any{{3, []any{5, -1}}, {2}, {1}, {-1}, {3}, {4}},
			output:   []any{nil, -1, 1, 1, 2, 3},
		},
		{
			commands: []string{"KthLargest", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add"},
			values:   [][]any{{7, []any{-10, 1, 3, 1, 4, 10, 3, 9, 4, 5, 1}}, {3}, {2}, {3}, {1}, {2}, {4}, {5}, {5}, {6}, {7}, {7}, {8}, {2}, {3}, {1}, {1}, {1}, {10}, {11}, {5}, {6}, {2}, {4}, {7}, {8}, {5}, {6}},
			output:   []any{nil, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		},
		{
			commands: []string{"KthLargest", "add", "add", "add", "add", "add"},
			values:   [][]any{{1, []any{}}, {-3}, {-2}, {-4}, {0}, {4}},
			output:   []any{nil, -3, -2, -2, 0, 4},
		},
	} {
		var kth KthLargest
		for index, command := range tc.commands {
			var output any = nil
			switch command {
			case "KthLargest":
				k := tc.values[index][0].(int)
				var src = tc.values[index][1].([]any)
				var nums = make([]int, 0, len(src))
				for _, num := range src {
					nums = append(nums, num.(int))
				}
				kth = NewKthLargest(k, nums)
			case "add":
				output = kth.Add(tc.values[index][0].(int))
			}
			assert.Equal(t, tc.output[index], output)
		}
	}
}
