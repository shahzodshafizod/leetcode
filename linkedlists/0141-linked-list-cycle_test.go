package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestHasCycle$
func TestHasCycle(t *testing.T) {
	for _, tc := range []struct {
		head *pkg.ListNode
		has  bool
	}{
		{head: pkg.MakeCycleLinkedList(1, 3, 2, 0, -4), has: true},
		{head: pkg.MakeCycleLinkedList(0, 1, 2), has: true},
		{head: pkg.MakeCycleLinkedList(-1, 1), has: false},
	} {
		has := hasCycle(tc.head)
		assert.Equal(t, tc.has, has)
	}
}
