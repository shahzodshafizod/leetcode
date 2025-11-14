package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMergeKLists$
func TestMergeKLists(t *testing.T) {
	for _, tc := range []struct {
		lists  []*pkg.ListNode
		sorted *pkg.ListNode
	}{
		{
			lists: []*pkg.ListNode{
				pkg.MakeLinkedList(1, 4, 5),
				pkg.MakeLinkedList(1, 3, 4),
				pkg.MakeLinkedList(2, 6),
			},
			sorted: pkg.MakeLinkedList(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			lists:  []*pkg.ListNode{},
			sorted: pkg.MakeLinkedList(),
		},
		{
			lists:  []*pkg.ListNode{pkg.MakeLinkedList()},
			sorted: pkg.MakeLinkedList(),
		},
		{
			lists: []*pkg.ListNode{
				pkg.MakeLinkedList(2),
				pkg.MakeLinkedList(),
				pkg.MakeLinkedList(-1),
			},
			sorted: pkg.MakeLinkedList(-1, 2),
		},
		{
			lists: []*pkg.ListNode{
				pkg.MakeLinkedList(-10, -9, -9, -3, -1, -1, 0),
				pkg.MakeLinkedList(-5),
				pkg.MakeLinkedList(4),
				pkg.MakeLinkedList(-8),
				pkg.MakeLinkedList(),
				pkg.MakeLinkedList(-9, -6, -5, -4, -2, 2, 3),
				pkg.MakeLinkedList(-3, -3, -2, -1, 0),
			},
			sorted: pkg.MakeLinkedList(-10, -9, -9, -9, -8, -6, -5, -5, -4, -3, -3, -3, -2, -2, -1, -1, -1, 0, 0, 2, 3, 4),
		},
		{
			lists: []*pkg.ListNode{
				pkg.MakeLinkedList(-8, -8, -3, -2, 0, 2, 2, 3, 3),
				pkg.MakeLinkedList(-5, -3, 1),
				pkg.MakeLinkedList(-9, -7, -1, 4),
				pkg.MakeLinkedList(-10, -4, -4, -4, 0, 3),
				pkg.MakeLinkedList(-2, 4),
				pkg.MakeLinkedList(-9, -6, -5, -5, -3, -3, -2, 2),
				pkg.MakeLinkedList(1, 3),
				pkg.MakeLinkedList(-8, -3, -2, 1, 3),
			},
			sorted: pkg.MakeLinkedList(-10, -9, -9, -8, -8, -8, -7, -6, -5, -5, -5, -4, -4, -4, -3, -3, -3, -3, -3, -2, -2, -2, -2, -1, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4),
		},
	} {
		sorted := mergeKLists(tc.lists)
		assert.Equal(t, tc.sorted, sorted)
	}
}
