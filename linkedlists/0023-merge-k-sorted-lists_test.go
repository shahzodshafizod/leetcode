package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMergeKLists$
func TestMergeKLists(t *testing.T) {
	for _, tc := range []struct {
		lists  []*design.ListNode
		sorted *design.ListNode
	}{
		{
			lists: []*design.ListNode{
				design.MakeLinkedList(1, 4, 5),
				design.MakeLinkedList(1, 3, 4),
				design.MakeLinkedList(2, 6),
			},
			sorted: design.MakeLinkedList(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			lists:  []*design.ListNode{},
			sorted: design.MakeLinkedList(),
		},
		{
			lists:  []*design.ListNode{design.MakeLinkedList()},
			sorted: design.MakeLinkedList(),
		},
		{
			lists: []*design.ListNode{
				design.MakeLinkedList(2),
				design.MakeLinkedList(),
				design.MakeLinkedList(-1),
			},
			sorted: design.MakeLinkedList(-1, 2),
		},
		{
			lists: []*design.ListNode{
				design.MakeLinkedList(-10, -9, -9, -3, -1, -1, 0),
				design.MakeLinkedList(-5),
				design.MakeLinkedList(4),
				design.MakeLinkedList(-8),
				design.MakeLinkedList(),
				design.MakeLinkedList(-9, -6, -5, -4, -2, 2, 3),
				design.MakeLinkedList(-3, -3, -2, -1, 0),
			},
			sorted: design.MakeLinkedList(-10, -9, -9, -9, -8, -6, -5, -5, -4, -3, -3, -3, -2, -2, -1, -1, -1, 0, 0, 2, 3, 4),
		},
		{
			lists: []*design.ListNode{
				design.MakeLinkedList(-8, -8, -3, -2, 0, 2, 2, 3, 3),
				design.MakeLinkedList(-5, -3, 1),
				design.MakeLinkedList(-9, -7, -1, 4),
				design.MakeLinkedList(-10, -4, -4, -4, 0, 3),
				design.MakeLinkedList(-2, 4),
				design.MakeLinkedList(-9, -6, -5, -5, -3, -3, -2, 2),
				design.MakeLinkedList(1, 3),
				design.MakeLinkedList(-8, -3, -2, 1, 3),
			},
			sorted: design.MakeLinkedList(-10, -9, -9, -8, -8, -8, -7, -6, -5, -5, -5, -4, -4, -4, -3, -3, -3, -3, -3, -2, -2, -2, -2, -1, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4),
		},
	} {
		sorted := mergeKLists(tc.lists)
		assert.Equal(t, tc.sorted, sorted)
	}
}
