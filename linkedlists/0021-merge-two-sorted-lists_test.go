package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMergeTwoLists$
func TestMergeTwoLists(t *testing.T) {
	for _, tc := range []struct {
		list1  *design.ListNode
		list2  *design.ListNode
		merged *design.ListNode
	}{
		{list1: design.MakeLinkedList(1, 2, 4), list2: design.MakeLinkedList(1, 3, 4), merged: design.MakeLinkedList(1, 1, 2, 3, 4, 4)},
		{list1: design.MakeLinkedList(), list2: design.MakeLinkedList(), merged: design.MakeLinkedList()},
		{list1: design.MakeLinkedList(), list2: design.MakeLinkedList(0), merged: design.MakeLinkedList(0)},
		{list1: design.MakeLinkedList(5, 10, 15, 40), list2: design.MakeLinkedList(2, 3, 20), merged: design.MakeLinkedList(2, 3, 5, 10, 15, 20, 40)},
	} {
		merged := mergeTwoLists(tc.list1, tc.list2)
		assert.Equal(t, tc.merged, merged)
	}
}
