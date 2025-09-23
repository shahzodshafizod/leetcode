package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestMedianFinder$
func TestMedianFinder(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			values:   [][]int{{}, {1}, {2}, {}, {3}, {}},
			output:   []any{nil, nil, nil, 1.5, nil, 2.0},
		},
		{
			commands: []string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
			values:   [][]int{{}, {6}, {}, {10}, {}, {2}, {}, {6}, {}, {5}, {}, {0}, {}, {6}, {}, {3}, {}, {1}, {}, {0}, {}, {0}, {}},
			output:   []any{nil, nil, 6.00000, nil, 8.00000, nil, 6.00000, nil, 6.00000, nil, 6.00000, nil, 5.50000, nil, 6.00000, nil, 5.50000, nil, 5.00000, nil, 4.00000, nil, 3.00000},
		},
		{
			commands: []string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
			values:   [][]int{{}, {1}, {}, {2}, {}, {3}, {}, {4}, {}, {5}, {}, {6}, {}, {7}, {}, {8}, {}, {9}, {}},
			output:   []any{nil, nil, 1.0, nil, 1.5, nil, 2.0, nil, 2.5, nil, 3.0, nil, 3.5, nil, 4.0, nil, 4.5, nil, 5.0},
		},
		{
			commands: []string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
			values:   [][]int{{}, {1}, {}, {1}, {}, {1}, {}, {1}, {}, {1}, {}, {1}, {}, {1}, {}, {1}, {}, {1}, {}},
			output:   []any{nil, nil, 1.0, nil, 1.0, nil, 1.0, nil, 1.0, nil, 1.0, nil, 1.0, nil, 1.0, nil, 1.0, nil, 1.0},
		},
	} {
		var medianFinder MedianFinder

		for index, command := range tc.commands {
			var output any

			switch command {
			case "MedianFinder":
				medianFinder = NewMedianFinder()
			case "addNum":
				medianFinder.AddNum(tc.values[index][0])
			case "findMedian":
				output = medianFinder.FindMedian()
			default:
			}

			assert.Equal(t, tc.output[index], output)
		}
	}
}
