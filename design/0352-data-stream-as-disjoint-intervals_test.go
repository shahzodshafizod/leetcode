package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestSummaryRanges$
func TestSummaryRanges(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		outputs  []any
	}{
		{
			commands: []string{"SummaryRanges", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals"},
			values:   [][]int{{}, {1}, {}, {3}, {}, {7}, {}, {2}, {}, {6}, {}},
			outputs:  []any{nil, nil, [][]int{{1, 1}}, nil, [][]int{{1, 1}, {3, 3}}, nil, [][]int{{1, 1}, {3, 3}, {7, 7}}, nil, [][]int{{1, 3}, {7, 7}}, nil, [][]int{{1, 3}, {6, 7}}},
		},
		{
			commands: []string{"SummaryRanges", "addNum", "getIntervals", "addNum", "getIntervals"},
			values:   [][]int{{}, {1}, {}, {0}, {}},
			outputs:  []any{nil, nil, [][]int{{1, 1}}, nil, [][]int{{0, 1}}},
		},
	} {
		var sr SummaryRanges

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "SummaryRanges":
				sr = NewSummaryRanges()
			case "addNum":
				sr.AddNum(tc.values[idx][0])
			case "getIntervals":
				output = sr.GetIntervals()
			default:
			}

			assert.Equal(t, tc.outputs[idx], output)
		}
	}
}
