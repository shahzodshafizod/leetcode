package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestRemoveZeroSumSublists$
func TestRemoveZeroSumSublists(t *testing.T) {
	for _, tc := range []struct {
		head    *pkg.ListNode
		newHead *pkg.ListNode
	}{
		{head: pkg.MakeLinkedList(1, 2, -3, 3, 1), newHead: pkg.MakeLinkedList(3, 1)},
		{head: pkg.MakeLinkedList(1, 2, 3, -3, 4), newHead: pkg.MakeLinkedList(1, 2, 4)},
		{head: pkg.MakeLinkedList(0, 0), newHead: pkg.MakeLinkedList()},
		{head: pkg.MakeLinkedList(1, 2, 3, -3, -2), newHead: pkg.MakeLinkedList(1)},
		{head: pkg.MakeLinkedList(-40, 40, 9, -2, 4), newHead: pkg.MakeLinkedList(9, -2, 4)},
		{head: pkg.MakeLinkedList(5, -3, -4, 1, 6, -2, -5), newHead: pkg.MakeLinkedList(5, -2, -5)},
	} {
		newHead := removeZeroSumSublists(tc.head)
		assert.Equal(t, newHead, newHead)
	}
}
