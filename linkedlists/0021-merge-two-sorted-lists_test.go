package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMergeTwoLists$
func TestMergeTwoLists(t *testing.T) {
	for _, tc := range []struct {
		list1  *pkg.ListNode[int]
		list2  *pkg.ListNode[int]
		merged *pkg.ListNode[int]
	}{
		{list1: pkg.MakeLinkedList(1, 2, 4), list2: pkg.MakeLinkedList(1, 3, 4), merged: pkg.MakeLinkedList(1, 1, 2, 3, 4, 4)},
		{list1: pkg.MakeLinkedList[int](), list2: pkg.MakeLinkedList[int](), merged: pkg.MakeLinkedList[int]()},
		{list1: pkg.MakeLinkedList[int](), list2: pkg.MakeLinkedList(0), merged: pkg.MakeLinkedList(0)},
		{list1: pkg.MakeLinkedList(5, 10, 15, 40), list2: pkg.MakeLinkedList(2, 3, 20), merged: pkg.MakeLinkedList(2, 3, 5, 10, 15, 20, 40)},
	} {
		merged := mergeTwoLists(tc.list1, tc.list2)
		assert.Equal(t, tc.merged, merged)
	}
}
