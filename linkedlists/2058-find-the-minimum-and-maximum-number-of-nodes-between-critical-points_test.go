package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestNodesBetweenCriticalPoints$
func TestNodesBetweenCriticalPoints(t *testing.T) {
	for _, tc := range []struct {
		head      *design.ListNode
		distances []int
	}{
		{head: design.MakeLinkedList(3, 1), distances: []int{-1, -1}},
		{head: design.MakeLinkedList(5, 3, 1, 2, 5, 1, 2), distances: []int{1, 3}},
		{head: design.MakeLinkedList(1, 3, 2, 2, 3, 2, 2, 2, 7), distances: []int{3, 3}},
	} {
		distances := nodesBetweenCriticalPoints(tc.head)
		assert.Equal(t, tc.distances, distances)
	}
}
