package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestReverseList$
func TestReverseList(t *testing.T) {
	for _, tc := range []struct {
		head     *pkg.ListNode
		reversed *pkg.ListNode
	}{
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5), reversed: pkg.MakeLinkedList(5, 4, 3, 2, 1)},
		{head: pkg.MakeLinkedList(3), reversed: pkg.MakeLinkedList(3)},
		{head: pkg.MakeLinkedList(), reversed: pkg.MakeLinkedList()},
		{head: pkg.MakeLinkedList(1, 2), reversed: pkg.MakeLinkedList(2, 1)},
	} {
		reversed := reverseList(tc.head)
		assert.Equal(t, tc.reversed, reversed)
	}
}
