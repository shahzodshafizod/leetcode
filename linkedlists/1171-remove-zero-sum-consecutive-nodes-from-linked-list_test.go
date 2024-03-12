package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestRemoveZeroSumSublists$
func TestRemoveZeroSumSublists(t *testing.T) {
	for _, tc := range []struct {
		head    *ListNode
		newHead *ListNode
	}{
		{head: makeLinkedList(1, 2, -3, 3, 1), newHead: makeLinkedList(3, 1)},
		{head: makeLinkedList(1, 2, 3, -3, 4), newHead: makeLinkedList(1, 2, 4)},
		{head: makeLinkedList(0, 0), newHead: makeLinkedList()},
		{head: makeLinkedList(1, 2, 3, -3, -2), newHead: makeLinkedList(1)},
		{head: makeLinkedList(-40, 40, 9, -2, 4), newHead: makeLinkedList(9, -2, 4)},
		{head: makeLinkedList(5, -3, -4, 1, 6, -2, -5), newHead: makeLinkedList(5, -2, -5)},
	} {
		newHead := removeZeroSumSublists(tc.head)
		assert.Equal(t, newHead, newHead)
	}
}
