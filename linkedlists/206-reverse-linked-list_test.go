package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dsa/linkedlists/ -run ^TestReverseList$
func TestReverseList(t *testing.T) {
	for _, data := range []struct {
		head     *ListNode
		reversed *ListNode
	}{
		{head: makeLinkedList(1, 2, 3, 4, 5), reversed: makeLinkedList(5, 4, 3, 2, 1)},
		{head: makeLinkedList(3), reversed: makeLinkedList(3)},
		{head: nil, reversed: nil},
		{head: makeLinkedList(), reversed: makeLinkedList()},
		{head: makeLinkedList(1, 2), reversed: makeLinkedList(2, 1)},
	} {
		reversed := reverseList(data.head)
		assert.Equal(t, data.reversed, reversed)
	}
}
