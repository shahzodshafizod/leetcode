package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestMiddleNode$
func TestMiddleNode(t *testing.T) {
	for _, tc := range []struct {
		head   *design.ListNode
		middle *design.ListNode
	}{
		{head: design.MakeLinkedList(1, 2, 3, 4, 5), middle: design.MakeLinkedList(3, 4, 5)},
		{head: design.MakeLinkedList(1, 2, 3, 4, 5, 6), middle: design.MakeLinkedList(4, 5, 6)},
	} {
		middle := middleNode(tc.head)
		assert.Equal(t, tc.middle, middle)
	}
}
