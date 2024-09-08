package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestSortList$
func TestSortList(t *testing.T) {
	for _, tc := range []struct {
		head   *design.ListNode
		sorted *design.ListNode
	}{
		{head: design.MakeLinkedList(), sorted: design.MakeLinkedList()},
		{head: design.MakeLinkedList(4), sorted: design.MakeLinkedList(4)},
		{head: design.MakeLinkedList(5, 2), sorted: design.MakeLinkedList(2, 5)},
		{head: design.MakeLinkedList(5, 2, 4), sorted: design.MakeLinkedList(2, 4, 5)},
		{head: design.MakeLinkedList(4, 2, 1, 3), sorted: design.MakeLinkedList(1, 2, 3, 4)},
		{head: design.MakeLinkedList(-1, 5, 3, 4, 0), sorted: design.MakeLinkedList(-1, 0, 3, 4, 5)},
	} {
		sorted := sortList(tc.head)
		assert.Equal(t, tc.sorted, sorted)
	}
}
