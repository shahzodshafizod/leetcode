package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./queues/ -run ^TestMyQueue$
func TestMyQueue(t *testing.T) {
	for _, data := range []struct {
		operations []string
		values     [][]int
		results    []any
	}{
		{
			operations: []string{"MyQueue", "push", "push", "peek", "pop", "empty"},
			values:     [][]int{{}, {1}, {2}, {}, {}, {}},
			results:    []any{nil, nil, nil, 1, 1, false},
		},
	} {
		var queue MyQueue
		for index, operation := range data.operations {
			var result any = nil
			switch operation {
			case "MyQueue":
				queue = Constructor()
			case "push":
				queue.Push(data.values[index][0])
			case "peek":
				result = queue.Peek()
			case "pop":
				result = queue.Pop()
			case "empty":
				result = queue.Empty()
			}
			assert.Equal(t, data.results[index], result)
		}
	}
}
