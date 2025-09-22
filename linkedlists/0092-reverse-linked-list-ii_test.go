package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestReverseBetween$
func TestReverseBetween(t *testing.T) {
	for _, tc := range []struct {
		head     *pkg.ListNode[int]
		left     int
		right    int
		reversed *pkg.ListNode[int]
	}{
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5), left: 2, right: 4, reversed: pkg.MakeLinkedList(1, 4, 3, 2, 5)},
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5), left: 1, right: 4, reversed: pkg.MakeLinkedList(4, 3, 2, 1, 5)},
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5, 6), left: 1, right: 6, reversed: pkg.MakeLinkedList(6, 5, 4, 3, 2, 1)},
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5, 6), left: 1, right: 2, reversed: pkg.MakeLinkedList(2, 1, 3, 4, 5, 6)},
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5, 6), left: 5, right: 6, reversed: pkg.MakeLinkedList(1, 2, 3, 4, 6, 5)},
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5, 6), left: 1, right: 1, reversed: pkg.MakeLinkedList(1, 2, 3, 4, 5, 6)},
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5, 6), left: 6, right: 6, reversed: pkg.MakeLinkedList(1, 2, 3, 4, 5, 6)},
		{head: pkg.MakeLinkedList(5), left: 1, right: 1, reversed: pkg.MakeLinkedList(5)},
		{head: pkg.MakeLinkedList[int](), left: 0, right: 0, reversed: pkg.MakeLinkedList[int]()},
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5, 6, 7), left: 3, right: 5, reversed: pkg.MakeLinkedList(1, 2, 5, 4, 3, 6, 7)},
	} {
		reversed := reverseBetween(tc.head, tc.left, tc.right)
		assert.Equal(t, tc.reversed, reversed)
	}
}
