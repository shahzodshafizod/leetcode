package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./pkg/ -run ^TestStack$
func TestStack(t *testing.T) {
	var stack Stack[string]

	for _, tc := range []struct {
		command string
		value   []string
		output  any
	}{
		{command: "Stack", value: []string{}, output: nil},
		{command: "Push", value: []string{"first"}, output: nil},
		{command: "Peek", value: []string{}, output: "first"},
		{command: "Size", value: []string{}, output: 1},
		{command: "Empty", value: []string{}, output: false},
		{command: "Pop", value: []string{}, output: "first"},
		{command: "Empty", value: []string{}, output: true},
		{command: "Push", value: []string{"second"}, output: nil},
		{command: "Push", value: []string{"third"}, output: nil},
		{command: "Peek", value: []string{}, output: "third"},
		{command: "Size", value: []string{}, output: 2},
		{command: "Empty", value: []string{}, output: false},
	} {
		var output any

		switch tc.command {
		case "Stack":
			stack = NewStack[string]()
		case "Push":
			stack.Push(tc.value[0])
		case "Pop":
			output = stack.Pop()
		case "Peek":
			output = stack.Top()
		case "Empty":
			output = stack.Empty()
		case "Size":
			output = stack.Size()
		}

		assert.Equal(t, tc.output, output)
	}
}
