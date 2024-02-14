package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMergeTwoLists$
func TestMergeTwoLists(t *testing.T) {
	for _, tc := range []struct {
		list1  *ListNode
		list2  *ListNode
		merged *ListNode
	}{
		{list1: makeLinkedList(1, 2, 4), list2: makeLinkedList(1, 3, 4), merged: makeLinkedList(1, 1, 2, 3, 4, 4)},
		{list1: makeLinkedList(), list2: makeLinkedList(), merged: makeLinkedList()},
		{list1: makeLinkedList(), list2: makeLinkedList(0), merged: makeLinkedList(0)},
		{list1: makeLinkedList(5, 10, 15, 40), list2: makeLinkedList(2, 3, 20), merged: makeLinkedList(2, 3, 5, 10, 15, 20, 40)},
	} {
		merged := mergeTwoLists(tc.list1, tc.list2)
		assert.Equal(t, tc.merged, merged)
	}
}
