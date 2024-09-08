package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestReverseBetween$
func TestReverseBetween(t *testing.T) {
	for _, tc := range []struct {
		head     *design.ListNode
		left     int
		right    int
		reversed *design.ListNode
	}{
		{head: design.MakeLinkedList(1, 2, 3, 4, 5), left: 2, right: 4, reversed: design.MakeLinkedList(1, 4, 3, 2, 5)},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5), left: 1, right: 4, reversed: design.MakeLinkedList(4, 3, 2, 1, 5)},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5, 6), left: 1, right: 6, reversed: design.MakeLinkedList(6, 5, 4, 3, 2, 1)},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5, 6), left: 1, right: 2, reversed: design.MakeLinkedList(2, 1, 3, 4, 5, 6)},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5, 6), left: 5, right: 6, reversed: design.MakeLinkedList(1, 2, 3, 4, 6, 5)},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5, 6), left: 1, right: 1, reversed: design.MakeLinkedList(1, 2, 3, 4, 5, 6)},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5, 6), left: 6, right: 6, reversed: design.MakeLinkedList(1, 2, 3, 4, 5, 6)},
		{head: design.MakeLinkedList(5), left: 1, right: 1, reversed: design.MakeLinkedList(5)},
		{head: design.MakeLinkedList(), left: 0, right: 0, reversed: design.MakeLinkedList()},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5, 6, 7), left: 3, right: 5, reversed: design.MakeLinkedList(1, 2, 5, 4, 3, 6, 7)},
	} {
		reversed := reverseBetween(tc.head, tc.left, tc.right)
		assert.Equal(t, tc.reversed, reversed)
	}
}
