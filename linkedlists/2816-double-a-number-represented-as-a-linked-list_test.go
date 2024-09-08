package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestDoubleIt$
func TestDoubleIt(t *testing.T) {
	for _, tc := range []struct {
		head    *design.ListNode
		newhead *design.ListNode
	}{
		{head: design.MakeLinkedList(1, 8, 9), newhead: design.MakeLinkedList(3, 7, 8)},
		{head: design.MakeLinkedList(9, 9, 9), newhead: design.MakeLinkedList(1, 9, 9, 8)},
	} {
		newhead := doubleIt(tc.head)
		assert.Equal(t, tc.newhead, newhead)
	}
}
