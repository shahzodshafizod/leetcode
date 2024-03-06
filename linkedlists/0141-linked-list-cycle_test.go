package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestHasCycle$
func TestHasCycle(t *testing.T) {
	const size = 3
	var heads [size]*ListNode
	var hases [size]bool
	heads[0], _ = makeCycleLinkedList(1, 3, 2, 0, -4)
	hases[0] = true
	heads[1], _ = makeCycleLinkedList(0, 1, 2)
	hases[1] = true
	heads[2], _ = makeCycleLinkedList(-1, 1)
	hases[2] = false

	for idx := 0; idx < size; idx++ {
		has := hasCycle(heads[idx])
		assert.Equal(t, hases[idx], has)
	}
}
