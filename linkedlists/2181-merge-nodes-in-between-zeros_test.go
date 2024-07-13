package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMergeNodes$
func TestMergeNodes(t *testing.T) {
	for _, tc := range []struct {
		head    *ListNode
		newHead *ListNode
	}{
		{head: makeLinkedList(0, 3, 1, 0, 4, 5, 2, 0), newHead: makeLinkedList(4, 11)},
		{head: makeLinkedList(0, 1, 0, 3, 0, 2, 2, 0), newHead: makeLinkedList(1, 3, 4)},
	} {
		newHead := mergeNodes(tc.head)
		assert.Equal(t, tc.newHead, newHead)
	}
}
