package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestReverseKGroup$
func TestReverseKGroup(t *testing.T) {
	for _, tc := range []struct {
		head     *design.ListNode
		k        int
		reversed *design.ListNode
	}{
		{head: design.MakeLinkedList(1, 2, 3, 4, 5), k: 2, reversed: design.MakeLinkedList(2, 1, 4, 3, 5)},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5), k: 3, reversed: design.MakeLinkedList(3, 2, 1, 4, 5)},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5), k: 4, reversed: design.MakeLinkedList(4, 3, 2, 1, 5)},
		{head: design.MakeLinkedList(), k: 2, reversed: design.MakeLinkedList()},
		{head: design.MakeLinkedList(1), k: 2, reversed: design.MakeLinkedList(1)},
	} {
		reversed := reverseKGroup(tc.head, tc.k)
		assert.Equal(t, tc.reversed, reversed)
	}
}
