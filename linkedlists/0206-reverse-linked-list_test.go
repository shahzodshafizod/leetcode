package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestReverseList$
func TestReverseList(t *testing.T) {
	for _, tc := range []struct {
		head     *design.ListNode
		reversed *design.ListNode
	}{
		{head: design.MakeLinkedList(1, 2, 3, 4, 5), reversed: design.MakeLinkedList(5, 4, 3, 2, 1)},
		{head: design.MakeLinkedList(3), reversed: design.MakeLinkedList(3)},
		{head: design.MakeLinkedList(), reversed: design.MakeLinkedList()},
		{head: design.MakeLinkedList(1, 2), reversed: design.MakeLinkedList(2, 1)},
	} {
		reversed := reverseList(tc.head)
		assert.Equal(t, tc.reversed, reversed)
	}
}
