package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestRemoveNodes$
func TestRemoveNodes(t *testing.T) {
	for _, tc := range []struct {
		head    *ListNode
		newhead *ListNode
	}{
		{head: makeLinkedList(5, 2, 13, 3, 8), newhead: makeLinkedList(13, 8)},
		{head: makeLinkedList(1, 1, 1, 1), newhead: makeLinkedList(1, 1, 1, 1)},
	} {
		newhead := removeNodes(tc.head)
		assert.Equal(t, tc.newhead, newhead)
	}
}
