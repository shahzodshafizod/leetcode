package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestSortList$
func TestSortList(t *testing.T) {
	for _, tc := range []struct {
		head   *ListNode
		sorted *ListNode
	}{
		{head: makeLinkedList(), sorted: makeLinkedList()},
		{head: makeLinkedList(4), sorted: makeLinkedList(4)},
		{head: makeLinkedList(5, 2), sorted: makeLinkedList(2, 5)},
		{head: makeLinkedList(5, 2, 4), sorted: makeLinkedList(2, 4, 5)},
		{head: makeLinkedList(4, 2, 1, 3), sorted: makeLinkedList(1, 2, 3, 4)},
		{head: makeLinkedList(-1, 5, 3, 4, 0), sorted: makeLinkedList(-1, 0, 3, 4, 5)},
	} {
		sorted := sortList(tc.head)
		assert.Equal(t, tc.sorted, sorted)
	}
}
