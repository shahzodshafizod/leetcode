package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestReorderList$
func TestReorderList(t *testing.T) {
	for _, tc := range []struct {
		head      *design.ListNode
		reordered *design.ListNode
	}{
		{head: design.MakeLinkedList(1, 2, 3, 4), reordered: design.MakeLinkedList(1, 4, 2, 3)},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9), reordered: design.MakeLinkedList(1, 9, 2, 8, 3, 7, 4, 6, 5)},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5), reordered: design.MakeLinkedList(1, 5, 2, 4, 3)},
		{head: design.MakeLinkedList(1), reordered: design.MakeLinkedList(1)},
	} {
		reorderList(tc.head)
		assert.Equal(t, tc.reordered, tc.head)
	}
}
