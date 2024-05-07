package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestDoubleIt$
func TestDoubleIt(t *testing.T) {
	for _, tc := range []struct {
		head    *ListNode
		newhead *ListNode
	}{
		{head: makeLinkedList(1, 8, 9), newhead: makeLinkedList(3, 7, 8)},
		{head: makeLinkedList(9, 9, 9), newhead: makeLinkedList(1, 9, 9, 8)},
	} {
		newhead := doubleIt(tc.head)
		assert.Equal(t, tc.newhead, newhead)
	}
}
