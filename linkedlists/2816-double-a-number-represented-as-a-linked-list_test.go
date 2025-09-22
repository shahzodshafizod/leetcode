package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestDoubleIt$
func TestDoubleIt(t *testing.T) {
	for _, tc := range []struct {
		head    *pkg.ListNode[int]
		newhead *pkg.ListNode[int]
	}{
		{head: pkg.MakeLinkedList(1, 8, 9), newhead: pkg.MakeLinkedList(3, 7, 8)},
		{head: pkg.MakeLinkedList(9, 9, 9), newhead: pkg.MakeLinkedList(1, 9, 9, 8)},
	} {
		newhead := doubleIt(tc.head)
		assert.Equal(t, tc.newhead, newhead)
	}
}
