package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestRemoveZeroSumSublists$
func TestRemoveZeroSumSublists(t *testing.T) {
	for _, tc := range []struct {
		head    *design.ListNode
		newHead *design.ListNode
	}{
		{head: design.MakeLinkedList(1, 2, -3, 3, 1), newHead: design.MakeLinkedList(3, 1)},
		{head: design.MakeLinkedList(1, 2, 3, -3, 4), newHead: design.MakeLinkedList(1, 2, 4)},
		{head: design.MakeLinkedList(0, 0), newHead: design.MakeLinkedList()},
		{head: design.MakeLinkedList(1, 2, 3, -3, -2), newHead: design.MakeLinkedList(1)},
		{head: design.MakeLinkedList(-40, 40, 9, -2, 4), newHead: design.MakeLinkedList(9, -2, 4)},
		{head: design.MakeLinkedList(5, -3, -4, 1, 6, -2, -5), newHead: design.MakeLinkedList(5, -2, -5)},
	} {
		newHead := removeZeroSumSublists(tc.head)
		assert.Equal(t, newHead, newHead)
	}
}
