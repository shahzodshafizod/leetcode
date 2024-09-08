package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestRemoveNthFromEnd$
func TestRemoveNthFromEnd(t *testing.T) {
	for _, tc := range []struct {
		head    *design.ListNode
		n       int
		newHead *design.ListNode
	}{
		{head: design.MakeLinkedList(1, 2, 3, 4, 5), n: 2, newHead: design.MakeLinkedList(1, 2, 3, 5)},
		{head: design.MakeLinkedList(1), n: 1, newHead: design.MakeLinkedList()},
		{head: design.MakeLinkedList(1, 2), n: 1, newHead: design.MakeLinkedList(1)},
	} {
		newHead := removeNthFromEnd(tc.head, tc.n)
		assert.Equal(t, tc.newHead, newHead)
	}
}
