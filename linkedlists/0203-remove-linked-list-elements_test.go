package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestRemoveElements$
func TestRemoveElements(t *testing.T) {
	for _, tc := range []struct {
		head    *pkg.ListNode[int]
		val     int
		newHead *pkg.ListNode[int]
	}{
		{head: pkg.MakeLinkedList(1, 2, 6, 3, 4, 5, 6), val: 6, newHead: pkg.MakeLinkedList(1, 2, 3, 4, 5)},
		{head: pkg.MakeLinkedList[int](), val: 1, newHead: pkg.MakeLinkedList[int]()},
		{head: pkg.MakeLinkedList(7, 7, 7, 7), val: 7, newHead: pkg.MakeLinkedList[int]()},
		{head: pkg.MakeLinkedList(1, 2, 2, 1), val: 2, newHead: pkg.MakeLinkedList(1, 1)},
		{head: pkg.MakeLinkedList(1), val: 2, newHead: pkg.MakeLinkedList(1)},
	} {
		newHead := removeElements(tc.head, tc.val)
		assert.Equal(t, tc.newHead, newHead)
	}
}
