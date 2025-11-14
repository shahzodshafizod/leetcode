package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMergeTwoLists$
func TestMergeTwoLists(t *testing.T) {
	for _, tc := range []struct {
		list1  *pkg.ListNode
		list2  *pkg.ListNode
		merged *pkg.ListNode
	}{
		{list1: pkg.MakeLinkedList(1, 2, 4), list2: pkg.MakeLinkedList(1, 3, 4), merged: pkg.MakeLinkedList(1, 1, 2, 3, 4, 4)},
		{list1: pkg.MakeLinkedList(), list2: pkg.MakeLinkedList(), merged: pkg.MakeLinkedList()},
		{list1: pkg.MakeLinkedList(), list2: pkg.MakeLinkedList(0), merged: pkg.MakeLinkedList(0)},
		{list1: pkg.MakeLinkedList(5, 10, 15, 40), list2: pkg.MakeLinkedList(2, 3, 20), merged: pkg.MakeLinkedList(2, 3, 5, 10, 15, 20, 40)},
	} {
		merged := mergeTwoLists(tc.list1, tc.list2)
		assert.Equal(t, tc.merged, merged)
	}
}
