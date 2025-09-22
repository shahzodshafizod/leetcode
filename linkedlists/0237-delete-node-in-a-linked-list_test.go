package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestDeleteNode$
func TestDeleteNode(t *testing.T) {
	for _, tc := range []struct {
		head     *pkg.ListNode[int]
		position int
		newHead  *pkg.ListNode[int]
	}{
		{head: pkg.MakeLinkedList(4, 5, 1, 9), position: 1, newHead: pkg.MakeLinkedList(4, 1, 9)},
		{head: pkg.MakeLinkedList(4, 5, 1, 9), position: 2, newHead: pkg.MakeLinkedList(4, 5, 9)},
	} {
		deleteNode(pkg.GetNode(tc.head, tc.position))
		assert.Equal(t, tc.newHead, tc.head)
	}
}
