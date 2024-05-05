package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestDeleteNode$
func TestDeleteNode(t *testing.T) {
	for _, tc := range []struct {
		head     *ListNode
		position int
		newHead  *ListNode
	}{
		{head: makeLinkedList(4, 5, 1, 9), position: 1, newHead: makeLinkedList(4, 1, 9)},
		{head: makeLinkedList(4, 5, 1, 9), position: 2, newHead: makeLinkedList(4, 5, 9)},
	} {
		deleteNode(getNode(tc.head, tc.position))
		assert.Equal(t, tc.newHead, tc.head)
	}
}
