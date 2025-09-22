package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestRemoveNodes$
func TestRemoveNodes(t *testing.T) {
	for _, tc := range []struct {
		head    *pkg.ListNode[int]
		newhead *pkg.ListNode[int]
	}{
		{head: pkg.MakeLinkedList(5, 2, 13, 3, 8), newhead: pkg.MakeLinkedList(13, 8)},
		{head: pkg.MakeLinkedList(1, 1, 1, 1), newhead: pkg.MakeLinkedList(1, 1, 1, 1)},
	} {
		newhead := removeNodes(tc.head)
		assert.Equal(t, tc.newhead, newhead)
	}
}
