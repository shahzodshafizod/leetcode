package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestDeleteDuplicates$
func TestDeleteDuplicates(t *testing.T) {
	for _, tc := range []struct {
		head    *pkg.ListNode
		cleaned *pkg.ListNode
	}{
		{head: pkg.MakeLinkedList(1, 1, 2), cleaned: pkg.MakeLinkedList(1, 2)},
		{head: pkg.MakeLinkedList(1, 1, 2, 3, 3), cleaned: pkg.MakeLinkedList(1, 2, 3)},
		{head: pkg.MakeLinkedList(0, 0, 0, 0, 0), cleaned: pkg.MakeLinkedList(0)},
	} {
		cleaned := deleteDuplicates(tc.head)
		assert.Equal(t, tc.cleaned, cleaned)
	}
}
