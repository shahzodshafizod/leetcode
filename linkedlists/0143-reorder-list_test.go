package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestReorderList$
func TestReorderList(t *testing.T) {
	for _, tc := range []struct {
		head      *ListNode
		reordered *ListNode
	}{
		{head: makeLinkedList(1, 2, 3, 4), reordered: makeLinkedList(1, 4, 2, 3)},
		{head: makeLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9), reordered: makeLinkedList(1, 9, 2, 8, 3, 7, 4, 6, 5)},
		{head: makeLinkedList(1, 2, 3, 4, 5), reordered: makeLinkedList(1, 5, 2, 4, 3)},
		{head: makeLinkedList(1), reordered: makeLinkedList(1)},
	} {
		reorderList(tc.head)
		assert.Equal(t, tc.reordered, tc.head)
	}
}
