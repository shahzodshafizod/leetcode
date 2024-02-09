package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./queues/ -run ^TestQueue$
func TestQueue(t *testing.T) {
	var queue Queue[int]
	for _, tc := range []struct {
		command string
		value   []int
		output  any
	}{
		{command: "Queue", value: []int{}, output: nil},
		{command: "Enqueue", value: []int{1}, output: nil},
		{command: "Enqueue", value: []int{2}, output: nil},
		{command: "Peek", value: []int{}, output: 1},
		{command: "Empty", value: []int{}, output: false},
		{command: "Size", value: []int{}, output: 2},
		{command: "Dequeue", value: []int{}, output: 1},
		{command: "Size", value: []int{}, output: 1},
		{command: "Dequeue", value: []int{}, output: 2},
		{command: "Empty", value: []int{}, output: true},
		{command: "Size", value: []int{}, output: 0},
		{command: "Enqueue", value: []int{3}, output: nil},
		{command: "Enqueue", value: []int{4}, output: nil},
		{command: "Size", value: []int{}, output: 2},
		{command: "Enqueue", value: []int{5}, output: nil},
		{command: "Size", value: []int{}, output: 3},
		{command: "Peek", value: []int{}, output: 3},
		{command: "Empty", value: []int{}, output: false},
		{command: "Dequeue", value: []int{}, output: 3},
		{command: "Dequeue", value: []int{}, output: 4},
		{command: "Dequeue", value: []int{}, output: 5},
		{command: "Empty", value: []int{}, output: true},
	} {
		var output any = nil
		switch tc.command {
		case "Queue":
			queue = NewQueue[int]()
		case "Enqueue":
			queue.Enqueue(tc.value[0])
		case "Dequeue":
			output = queue.Dequeue()
		case "Peek":
			output = queue.Peek()
		case "Empty":
			output = queue.Empty()
		case "Size":
			output = queue.Size()
		}
		assert.Equal(t, tc.output, output)
	}
}
