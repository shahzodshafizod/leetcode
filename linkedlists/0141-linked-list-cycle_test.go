package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestHasCycle$
func TestHasCycle(t *testing.T) {
	for _, tc := range []struct {
		head *design.ListNode
		has  bool
	}{
		{head: design.MakeCycleLinkedList(1, 3, 2, 0, -4), has: true},
		{head: design.MakeCycleLinkedList(0, 1, 2), has: true},
		{head: design.MakeCycleLinkedList(-1, 1), has: false},
	} {
		has := hasCycle(tc.head)
		assert.Equal(t, tc.has, has)
	}
}
