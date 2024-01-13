package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestDetectCycle$
func TestDetectCycle(t *testing.T) {
	// head, cyclePoint := makeCycleLinkedList(2, 1, 2, 3, 4, 5, 6, 7, 8)
	// head, cyclePoint := makeCycleLinkedList(7, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	// head, cyclePoint := makeCycleLinkedList(1, 3, 2, 0, -4)
	// head, cyclePoint := makeCycleLinkedList(0, 1, 2)
	// head, cyclePoint := makeCycleLinkedList(-1, 1, 2)
	// head, cyclePoint := makeCycleLinkedList(-1, 1)
	head, cyclePoint := makeCycleLinkedList(-1)

	cycle := detectCycle(head)
	assert.EqualValues(t, cyclePoint, cycle)
}
