package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestIsPalindrome$
func TestIsPalindrome(t *testing.T) {
	for _, tc := range []struct {
		head *design.ListNode
		is   bool
	}{
		{head: design.MakeLinkedList(1, 2, 3, 2, 1), is: true},
		{head: design.MakeLinkedList(1, 2, 2, 1), is: true},
		{head: design.MakeLinkedList(1, 2), is: false},
		{head: design.MakeLinkedList(1), is: true},
	} {
		is := isPalindrome(tc.head)
		assert.Equal(t, tc.is, is)
	}
}
