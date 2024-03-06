package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestDetectCycle$
func TestDetectCycle(t *testing.T) {
	const size = 7
	var heads [size]*ListNode
	var cyclePoints [size]*ListNode
	heads[0], cyclePoints[0] = makeCycleLinkedList(2, 1, 2, 3, 4, 5, 6, 7, 8)
	heads[1], cyclePoints[1] = makeCycleLinkedList(7, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	heads[2], cyclePoints[2] = makeCycleLinkedList(1, 3, 2, 0, -4)
	heads[3], cyclePoints[3] = makeCycleLinkedList(0, 1, 2)
	heads[4], cyclePoints[4] = makeCycleLinkedList(-1, 1, 2)
	heads[5], cyclePoints[5] = makeCycleLinkedList(-1, 1)
	heads[6], cyclePoints[6] = makeCycleLinkedList(-1)

	for idx := 0; idx < size; idx++ {
		cyclePoint := detectCycle(heads[idx])
		assert.EqualValues(t, cyclePoints[idx], cyclePoint)
	}
}
