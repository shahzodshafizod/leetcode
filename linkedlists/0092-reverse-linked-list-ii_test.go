package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestReverseBetween$
func TestReverseBetween(t *testing.T) {
	for _, data := range []struct {
		head     *ListNode
		left     int
		right    int
		reversed *ListNode
	}{
		{head: makeLinkedList(1, 2, 3, 4, 5), left: 2, right: 4, reversed: makeLinkedList(1, 4, 3, 2, 5)},
		{head: makeLinkedList(1, 2, 3, 4, 5), left: 1, right: 4, reversed: makeLinkedList(4, 3, 2, 1, 5)},
		{head: makeLinkedList(1, 2, 3, 4, 5, 6), left: 1, right: 6, reversed: makeLinkedList(6, 5, 4, 3, 2, 1)},
		{head: makeLinkedList(1, 2, 3, 4, 5, 6), left: 1, right: 2, reversed: makeLinkedList(2, 1, 3, 4, 5, 6)},
		{head: makeLinkedList(1, 2, 3, 4, 5, 6), left: 5, right: 6, reversed: makeLinkedList(1, 2, 3, 4, 6, 5)},
		{head: makeLinkedList(1, 2, 3, 4, 5, 6), left: 1, right: 1, reversed: makeLinkedList(1, 2, 3, 4, 5, 6)},
		{head: makeLinkedList(1, 2, 3, 4, 5, 6), left: 6, right: 6, reversed: makeLinkedList(1, 2, 3, 4, 5, 6)},
		{head: makeLinkedList(5), left: 1, right: 1, reversed: makeLinkedList(5)},
		{head: makeLinkedList(), left: 0, right: 0, reversed: makeLinkedList()},
		{head: makeLinkedList(1, 2, 3, 4, 5, 6, 7), left: 3, right: 5, reversed: makeLinkedList(1, 2, 5, 4, 3, 6, 7)},
	} {
		reversed := reverseBetween(data.head, data.left, data.right)
		assert.Equal(t, data.reversed, reversed)
	}
}
