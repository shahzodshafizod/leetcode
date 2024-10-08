package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestSortList$
func TestSortList(t *testing.T) {
	for _, tc := range []struct {
		head   *pkg.ListNode
		sorted *pkg.ListNode
	}{
		{head: pkg.MakeLinkedList(), sorted: pkg.MakeLinkedList()},
		{head: pkg.MakeLinkedList(4), sorted: pkg.MakeLinkedList(4)},
		{head: pkg.MakeLinkedList(5, 2), sorted: pkg.MakeLinkedList(2, 5)},
		{head: pkg.MakeLinkedList(5, 2, 4), sorted: pkg.MakeLinkedList(2, 4, 5)},
		{head: pkg.MakeLinkedList(4, 2, 1, 3), sorted: pkg.MakeLinkedList(1, 2, 3, 4)},
		{head: pkg.MakeLinkedList(-1, 5, 3, 4, 0), sorted: pkg.MakeLinkedList(-1, 0, 3, 4, 5)},
	} {
		sorted := sortList(tc.head)
		assert.Equal(t, tc.sorted, sorted)
	}
}
