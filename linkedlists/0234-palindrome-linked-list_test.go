package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestIsPalindrome$
func TestIsPalindrome(t *testing.T) {
	for _, tc := range []struct {
		head *ListNode
		is   bool
	}{
		{head: makeLinkedList(1, 2, 3, 2, 1), is: true},
		{head: makeLinkedList(1, 2, 2, 1), is: true},
		{head: makeLinkedList(1, 2), is: false},
		{head: makeLinkedList(1), is: true},
	} {
		is := isPalindrome(tc.head)
		assert.Equal(t, tc.is, is)
	}
}
