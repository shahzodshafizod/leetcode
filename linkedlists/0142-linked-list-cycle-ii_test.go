package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestDetectCycle$
func TestDetectCycle(t *testing.T) {
	for _, tc := range []struct {
		head     *ListNode
		position int
	}{
		{head: makeCycleLinkedList(2, 1, 2, 3, 4, 5, 6, 7, 8), position: 2},
		{head: makeCycleLinkedList(7, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12), position: 7},
		{head: makeCycleLinkedList(1, 3, 2, 0, -4), position: 1},
		{head: makeCycleLinkedList(0, 1, 2), position: 0},
		{head: makeCycleLinkedList(-1, 1, 2), position: -1},
		{head: makeCycleLinkedList(-1, 1), position: -1},
		{head: makeCycleLinkedList(-1), position: -1},
	} {
		cyclePoint := detectCycle(tc.head)
		assert.EqualValues(t, getNode(tc.head, tc.position), cyclePoint)
	}
}
