package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestFreqStack$
func TestFreqStack(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"FreqStack", "push", "push", "push", "push", "push", "push", "pop", "pop", "pop", "pop"},
			values:   [][]int{{}, {5}, {7}, {5}, {7}, {4}, {5}, {}, {}, {}, {}},
			output:   []any{nil, nil, nil, nil, nil, nil, nil, 5, 7, 5, 4},
		},
		{
			commands: []string{"FreqStack", "push", "push", "push", "push", "push", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "push", "pop", "pop", "pop", "pop", "pop", "pop"},
			values:   [][]int{{}, {4}, {0}, {9}, {3}, {4}, {2}, {}, {6}, {}, {1}, {}, {1}, {}, {4}, {}, {}, {}, {}, {}, {}},
			output:   []any{nil, nil, nil, nil, nil, nil, nil, 4, nil, 6, nil, 1, nil, 1, nil, 4, 2, 3, 9, 0, 4},
		},
	} {
		var f FreqStack

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "FreqStack":
				f = NewFreqStack()
			case "push":
				f.Push(tc.values[idx][0])
			case "pop":
				output = f.Pop()
			default:
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
