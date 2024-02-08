package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestReverseKGroup$
func TestReverseKGroup(t *testing.T) {
	for _, tc := range []struct {
		head     *ListNode
		k        int
		reversed *ListNode
	}{
		{head: makeLinkedList(1, 2, 3, 4, 5), k: 2, reversed: makeLinkedList(2, 1, 4, 3, 5)},
		{head: makeLinkedList(1, 2, 3, 4, 5), k: 3, reversed: makeLinkedList(3, 2, 1, 4, 5)},
		{head: makeLinkedList(1, 2, 3, 4, 5), k: 4, reversed: makeLinkedList(4, 3, 2, 1, 5)},
		{head: makeLinkedList(), k: 2, reversed: makeLinkedList()},
		{head: makeLinkedList(1), k: 2, reversed: makeLinkedList(1)},
	} {
		reversed := reverseKGroup(tc.head, tc.k)
		assert.Equal(t, tc.reversed, reversed)
	}
}
