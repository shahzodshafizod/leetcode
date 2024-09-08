package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMergeInBetween$
func TestMergeInBetween(t *testing.T) {
	for _, tc := range []struct {
		list1  *design.ListNode
		a      int
		b      int
		list2  *design.ListNode
		merged *design.ListNode
	}{
		{
			list1:  design.MakeLinkedList(10, 1, 13, 6, 9, 5),
			a:      3,
			b:      4,
			list2:  design.MakeLinkedList(1000000, 1000001, 1000002),
			merged: design.MakeLinkedList(10, 1, 13, 1000000, 1000001, 1000002, 5),
		},
		{
			list1:  design.MakeLinkedList(0, 1, 2, 3, 4, 5, 6),
			a:      2,
			b:      5,
			list2:  design.MakeLinkedList(1000000, 1000001, 1000002, 1000003, 1000004),
			merged: design.MakeLinkedList(0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6),
		},
		{
			list1:  design.MakeLinkedList(0, 3, 2, 1, 4, 5),
			a:      3,
			b:      4,
			list2:  design.MakeLinkedList(1000000, 1000001, 1000002),
			merged: design.MakeLinkedList(0, 3, 2, 1000000, 1000001, 1000002, 5),
		},
	} {
		merged := mergeInBetween(tc.list1, tc.a, tc.b, tc.list2)
		assert.Equal(t, tc.merged, merged)
	}
}
