package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestTaskManager$
func TestTaskManager(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   []any
		output   []any
	}{
		{
			commands: []string{"TaskManager", "add", "edit", "execTop", "rmv", "add", "execTop"},
			values:   []any{[][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}}, []int{4, 104, 5}, []int{102, 8}, []int{}, []int{101}, []int{5, 105, 15}, []int{}},
			output:   []any{nil, nil, nil, 3, nil, nil, 5},
		},
		{
			commands: []string{"TaskManager", "edit", "execTop", "edit", "rmv", "rmv", "execTop", "execTop"},
			values:   []any{[][]int{{7, 13, 16}, {3, 11, 41}, {7, 18, 40}}, []int{13, 0}, []int{}, []int{13, 19}, []int{18}, []int{13}, []int{}, []int{}},
			output:   []any{nil, nil, 3, nil, nil, nil, -1, -1},
		},
	} {
		var manager TaskManager

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "TaskManager":
				tasks, ok := tc.values[idx].([][]int)
				_ = ok
				manager = NewTaskManager(tasks)
			case "add":
				values, ok := tc.values[idx].([]int)
				_ = ok

				manager.Add(values[0], values[1], values[2])
			case "edit":
				values, ok := tc.values[idx].([]int)
				_ = ok

				manager.Edit(values[0], values[1])
			case "rmv":
				values, ok := tc.values[idx].([]int)
				_ = ok

				manager.Rmv(values[0])
			case "execTop":
				output = manager.ExecTop()
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
