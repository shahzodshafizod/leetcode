package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestRemoveNthFromEnd$
func TestRemoveNthFromEnd(t *testing.T) {
	for _, tc := range []struct {
		head    *ListNode
		n       int
		newHead *ListNode
	}{
		{head: makeLinkedList(1, 2, 3, 4, 5), n: 2, newHead: makeLinkedList(1, 2, 3, 5)},
		{head: makeLinkedList(1), n: 1, newHead: makeLinkedList()},
		{head: makeLinkedList(1, 2), n: 1, newHead: makeLinkedList(1)},
	} {
		newHead := removeNthFromEnd(tc.head, tc.n)
		assert.Equal(t, tc.newHead, newHead)
	}
}
