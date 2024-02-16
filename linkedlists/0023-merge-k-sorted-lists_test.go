package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMergeKLists$
func TestMergeKLists(t *testing.T) {
	for _, tc := range []struct {
		lists  []*ListNode
		sorted *ListNode
	}{
		{
			lists: []*ListNode{
				makeLinkedList(1, 4, 5),
				makeLinkedList(1, 3, 4),
				makeLinkedList(2, 6),
			},
			sorted: makeLinkedList(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			lists:  []*ListNode{},
			sorted: makeLinkedList(),
		},
		{
			lists:  []*ListNode{makeLinkedList()},
			sorted: makeLinkedList(),
		},
		{
			lists: []*ListNode{
				makeLinkedList(2),
				makeLinkedList(),
				makeLinkedList(-1),
			},
			sorted: makeLinkedList(-1, 2),
		},
		{
			lists: []*ListNode{
				makeLinkedList(-10, -9, -9, -3, -1, -1, 0),
				makeLinkedList(-5),
				makeLinkedList(4),
				makeLinkedList(-8),
				makeLinkedList(),
				makeLinkedList(-9, -6, -5, -4, -2, 2, 3),
				makeLinkedList(-3, -3, -2, -1, 0),
			},
			sorted: makeLinkedList(-10, -9, -9, -9, -8, -6, -5, -5, -4, -3, -3, -3, -2, -2, -1, -1, -1, 0, 0, 2, 3, 4),
		},
		{
			lists: []*ListNode{
				makeLinkedList(-8, -8, -3, -2, 0, 2, 2, 3, 3),
				makeLinkedList(-5, -3, 1),
				makeLinkedList(-9, -7, -1, 4),
				makeLinkedList(-10, -4, -4, -4, 0, 3),
				makeLinkedList(-2, 4),
				makeLinkedList(-9, -6, -5, -5, -3, -3, -2, 2),
				makeLinkedList(1, 3),
				makeLinkedList(-8, -3, -2, 1, 3),
			},
			sorted: makeLinkedList(-10, -9, -9, -8, -8, -8, -7, -6, -5, -5, -5, -4, -4, -4, -3, -3, -3, -3, -3, -2, -2, -2, -2, -1, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4),
		},
	} {
		sorted := mergeKLists(tc.lists)
		assert.Equal(t, tc.sorted, sorted)
	}
}
