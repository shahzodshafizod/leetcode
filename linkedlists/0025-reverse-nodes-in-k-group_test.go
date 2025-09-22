package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestReverseKGroup$
func TestReverseKGroup(t *testing.T) {
	for _, tc := range []struct {
		head     *pkg.ListNode[int]
		k        int
		reversed *pkg.ListNode[int]
	}{
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5), k: 2, reversed: pkg.MakeLinkedList(2, 1, 4, 3, 5)},
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5), k: 3, reversed: pkg.MakeLinkedList(3, 2, 1, 4, 5)},
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5), k: 4, reversed: pkg.MakeLinkedList(4, 3, 2, 1, 5)},
		{head: pkg.MakeLinkedList[int](), k: 2, reversed: pkg.MakeLinkedList[int]()},
		{head: pkg.MakeLinkedList(1), k: 2, reversed: pkg.MakeLinkedList(1)},
	} {
		reversed := reverseKGroup(tc.head, tc.k)
		assert.Equal(t, tc.reversed, reversed)
	}
}
