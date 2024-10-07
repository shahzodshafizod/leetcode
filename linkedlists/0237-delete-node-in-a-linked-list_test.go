package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestDeleteNode$
func TestDeleteNode(t *testing.T) {
	for _, tc := range []struct {
		head     *design.ListNode
		position int
		newHead  *design.ListNode
	}{
		{head: design.MakeLinkedList(4, 5, 1, 9), position: 1, newHead: design.MakeLinkedList(4, 1, 9)},
		{head: design.MakeLinkedList(4, 5, 1, 9), position: 2, newHead: design.MakeLinkedList(4, 5, 9)},
	} {
		deleteNode(design.GetNode(tc.head, tc.position))
		assert.Equal(t, tc.newHead, tc.head)
	}
}
