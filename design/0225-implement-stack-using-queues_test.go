package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestMyStack$
func TestMyStack(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"MyStack", "push", "push", "top", "pop", "empty"},
			values:   [][]int{{}, {1}, {2}, {}, {}, {}},
			output:   []any{nil, nil, nil, 2, 2, false},
		},
		{
			commands: []string{"MyStack", "push", "push", "top", "pop", "pop", "empty"},
			values:   [][]int{{}, {1}, {2}, {}, {}, {}, {}},
			output:   []any{nil, nil, nil, 2, 2, 1, true},
		},
	} {
		var stack MyStack

		for index, command := range tc.commands {
			var output any = nil

			switch command {
			case "MyStack":
				stack = NewMyStack()
			case "push":
				stack.Push(tc.values[index][0])
			case "pop":
				output = stack.Pop()
			case "top":
				output = stack.Top()
			case "empty":
				output = stack.Empty()
			}

			assert.Equal(t, tc.output[index], output)
		}
	}
}
