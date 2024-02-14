package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestMyLinkedList$
func TestMyLinkedList(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"},
			values:   [][]int{{}, {1}, {3}, {1, 2}, {1}, {1}, {1}},
			output:   []any{nil, nil, nil, nil, 2, nil, 3},
		},
		{
			commands: []string{"MyLinkedList", "addAtHead", "addAtHead", "addAtHead", "addAtIndex", "deleteAtIndex", "addAtHead", "addAtTail", "get", "addAtHead", "addAtIndex", "addAtHead"},
			values:   [][]int{{}, {7}, {2}, {1}, {3, 0}, {2}, {6}, {4}, {4}, {4}, {5, 0}, {6}},
			output:   []any{nil, nil, nil, nil, nil, nil, nil, nil, 4, nil, nil, nil},
		},
		{
			commands: []string{"MyLinkedList", "addAtHead", "addAtIndex", "get", "get", "get"},
			values:   [][]int{{}, {1}, {1, 2}, {1}, {0}, {2}},
			output:   []any{nil, nil, nil, 2, 1, -1},
		},
		{
			commands: []string{"MyLinkedList", "addAtHead", "get", "addAtHead", "addAtHead", "deleteAtIndex", "addAtHead", "get", "get", "get", "addAtHead", "deleteAtIndex"},
			values:   [][]int{{}, {4}, {1}, {1}, {5}, {3}, {7}, {3}, {3}, {3}, {1}, {4}},
			output:   []any{nil, nil, -1, nil, nil, nil, nil, 4, 4, 4, nil, nil},
		},
		{
			commands: []string{"MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"},
			values:   [][]int{{}, {1}, {3}, {1, 2}, {1}, {0}, {0}},
			output:   []any{nil, nil, nil, nil, 2, nil, 2},
		},
		{
			commands: []string{"MyLinkedList", "addAtIndex", "addAtIndex", "addAtIndex", "get"},
			values:   [][]int{{}, {0, 10}, {0, 20}, {1, 30}, {0}},
			output:   []any{nil, nil, nil, nil, 20},
		},
	} {
		var list MyLinkedList
		for index, command := range tc.commands {
			var output any = nil
			switch command {
			case "MyLinkedList":
				list = NewMyLinkedList()
			case "get":
				output = list.Get(tc.values[index][0])
			case "addAtHead":
				list.AddAtHead(tc.values[index][0])
			case "addAtTail":
				list.AddAtTail(tc.values[index][0])
			case "addAtIndex":
				list.AddAtIndex(tc.values[index][0], tc.values[index][1])
			case "deleteAtIndex":
				list.DeleteAtIndex(tc.values[index][0])
			}
			assert.Equal(t, tc.output[index], output)
		}
	}
}
