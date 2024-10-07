package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestRemoveNodes$
func TestRemoveNodes(t *testing.T) {
	for _, tc := range []struct {
		head    *design.ListNode
		newhead *design.ListNode
	}{
		{head: design.MakeLinkedList(5, 2, 13, 3, 8), newhead: design.MakeLinkedList(13, 8)},
		{head: design.MakeLinkedList(1, 1, 1, 1), newhead: design.MakeLinkedList(1, 1, 1, 1)},
	} {
		newhead := removeNodes(tc.head)
		assert.Equal(t, tc.newhead, newhead)
	}
}
