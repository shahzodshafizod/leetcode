package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMiddleNode$
func TestMiddleNode(t *testing.T) {
	for _, tc := range []struct {
		head   *pkg.ListNode
		middle *pkg.ListNode
	}{
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5), middle: pkg.MakeLinkedList(3, 4, 5)},
		{head: pkg.MakeLinkedList(1, 2, 3, 4, 5, 6), middle: pkg.MakeLinkedList(4, 5, 6)},
	} {
		middle := middleNode(tc.head)
		assert.Equal(t, tc.middle, middle)
	}
}
