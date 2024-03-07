package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMiddleNode$
func TestMiddleNode(t *testing.T) {
	for _, tc := range []struct {
		head   *ListNode
		middle *ListNode
	}{
		{head: makeLinkedList(1, 2, 3, 4, 5), middle: makeLinkedList(3, 4, 5)},
		{head: makeLinkedList(1, 2, 3, 4, 5, 6), middle: makeLinkedList(4, 5, 6)},
	} {
		middle := middleNode(tc.head)
		assert.Equal(t, tc.middle, middle)
	}
}
