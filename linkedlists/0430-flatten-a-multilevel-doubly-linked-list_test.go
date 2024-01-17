package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestFlatten$
func TestFlatten(t *testing.T) {
	for _, data := range []struct {
		root          *Node
		flattenedRoot *Node
	}{
		{root: makeDoublyLinkedList(1, 2, []any{7, 8, []any{10, 11}, 9}, 3, 4, 5, []any{12, 13}, 6),
			flattenedRoot: makeDoublyLinkedList(1, 2, 7, 8, 10, 11, 9, 3, 4, 5, 12, 13, 6)},
		{root: makeDoublyLinkedList(1, 2, 3, []any{7, 8, []any{11, 12}, 9, 10}, 4, 5, 6),
			flattenedRoot: makeDoublyLinkedList(1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6)},
		{root: makeDoublyLinkedList(1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6),
			flattenedRoot: makeDoublyLinkedList(1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6)},
		{root: makeDoublyLinkedList(1, []any{2, []any{3}}),
			flattenedRoot: makeDoublyLinkedList(1, 2, 3)},
		{root: makeDoublyLinkedList(), flattenedRoot: makeDoublyLinkedList()},
	} {
		flattenedRoot := flatten(data.root)
		assert.Equal(t, data.flattenedRoot, flattenedRoot)
	}
}
