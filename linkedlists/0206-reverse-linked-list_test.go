package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestReverseList$
func TestReverseList(t *testing.T) {
	for _, tc := range []struct {
		head     *ListNode
		reversed *ListNode
	}{
		{head: makeLinkedList(1, 2, 3, 4, 5), reversed: makeLinkedList(5, 4, 3, 2, 1)},
		{head: makeLinkedList(3), reversed: makeLinkedList(3)},
		{head: makeLinkedList(), reversed: makeLinkedList()},
		{head: makeLinkedList(1, 2), reversed: makeLinkedList(2, 1)},
	} {
		reversed := reverseList(tc.head)
		assert.Equal(t, tc.reversed, reversed)
	}
}
