package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestRemoveNthFromEnd$
func TestRemoveNthFromEnd(t *testing.T) {
	for _, tc := range []struct {
		head    *pkg.ListNode[int]
		n       int
		newHead *pkg.ListNode[int]
	}{
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5), n: 2, newHead: pkg.MakeLinkedList(1, 2, 3, 5)},
		{head: pkg.MakeLinkedList(1), n: 1, newHead: pkg.MakeLinkedList[int]()},
		{head: pkg.MakeLinkedList(1, 2), n: 1, newHead: pkg.MakeLinkedList(1)},
	} {
		newHead := removeNthFromEnd(tc.head, tc.n)
		assert.Equal(t, tc.newHead, newHead)
	}
}
