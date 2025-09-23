package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestNumberContainers$
func TestNumberContainers(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"},
			values:   [][]int{{}, {10}, {2, 10}, {1, 10}, {3, 10}, {5, 10}, {10}, {1, 20}, {10}},
			output:   []any{nil, -1, nil, nil, nil, nil, 1, nil, 2},
		},
	} {
		var container NumberContainers

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "NumberContainers":
				container = NewNumberContainers()
			case "change":
				container.Change(tc.values[idx][0], tc.values[idx][1])
			case "find":
				output = container.Find(tc.values[idx][0])
			default:
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
