package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestHasCycle$
func TestHasCycle(t *testing.T) {
	for _, tc := range []struct {
		head *ListNode
		has  bool
	}{
		{head: makeCycleLinkedList(1, 3, 2, 0, -4), has: true},
		{head: makeCycleLinkedList(0, 1, 2), has: true},
		{head: makeCycleLinkedList(-1, 1), has: false},
	} {
		has := hasCycle(tc.head)
		assert.Equal(t, tc.has, has)
	}
}
