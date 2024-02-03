package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestMyQueue$
func TestMyQueue(t *testing.T) {
	for _, tc := range []struct {
		operations []string
		values     [][]int
		output     []any
	}{
		{
			operations: []string{"MyQueue", "push", "push", "peek", "pop", "empty"},
			values:     [][]int{{}, {1}, {2}, {}, {}, {}},
			output:     []any{nil, nil, nil, 1, 1, false},
		},
	} {
		var queue MyQueue
		for index, operation := range tc.operations {
			var output any = nil
			switch operation {
			case "MyQueue":
				queue = NewMyQueue()
			case "push":
				queue.Push(tc.values[index][0])
			case "peek":
				output = queue.Peek()
			case "pop":
				output = queue.Pop()
			case "empty":
				output = queue.Empty()
			}
			assert.Equal(t, tc.output[index], output)
		}
	}
}
