package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestDetectCycle$
func TestDetectCycle(t *testing.T) {
	for _, tc := range []struct {
		head     *design.ListNode
		position int
	}{
		{head: design.MakeCycleLinkedList(2, 1, 2, 3, 4, 5, 6, 7, 8), position: 2},
		{head: design.MakeCycleLinkedList(7, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12), position: 7},
		{head: design.MakeCycleLinkedList(1, 3, 2, 0, -4), position: 1},
		{head: design.MakeCycleLinkedList(0, 1, 2), position: 0},
		{head: design.MakeCycleLinkedList(-1, 1, 2), position: -1},
		{head: design.MakeCycleLinkedList(-1, 1), position: -1},
		{head: design.MakeCycleLinkedList(-1), position: -1},
	} {
		cyclePoint := detectCycle(tc.head)
		assert.EqualValues(t, design.GetNode(tc.head, tc.position), cyclePoint)
	}
}
