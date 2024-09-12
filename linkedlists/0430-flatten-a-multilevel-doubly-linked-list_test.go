package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestFlatten$
func TestFlatten(t *testing.T) {
	for _, tc := range []struct {
		root          *design.DListNode
		flattenedRoot *design.DListNode
	}{
		{root: design.MakeDoublyLinkedList(1, 2, []any{7, 8, []any{10, 11}, 9}, 3, 4, 5, []any{12, 13}, 6),
			flattenedRoot: design.MakeDoublyLinkedList(1, 2, 7, 8, 10, 11, 9, 3, 4, 5, 12, 13, 6)},
		{root: design.MakeDoublyLinkedList(1, 2, 3, []any{7, 8, []any{11, 12}, 9, 10}, 4, 5, 6),
			flattenedRoot: design.MakeDoublyLinkedList(1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6)},
		{root: design.MakeDoublyLinkedList(1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6),
			flattenedRoot: design.MakeDoublyLinkedList(1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6)},
		{root: design.MakeDoublyLinkedList(1, []any{2, []any{3}}),
			flattenedRoot: design.MakeDoublyLinkedList(1, 2, 3)},
		{root: design.MakeDoublyLinkedList(), flattenedRoot: design.MakeDoublyLinkedList()},
	} {
		flattenedRoot := flatten(tc.root)
		assert.Equal(t, tc.flattenedRoot, flattenedRoot)
	}
}
