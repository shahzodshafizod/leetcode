package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestNodesBetweenCriticalPoints$
func TestNodesBetweenCriticalPoints(t *testing.T) {
	for _, tc := range []struct {
		head      *pkg.ListNode
		distances []int
	}{
		{head: pkg.MakeLinkedList(3, 1), distances: []int{-1, -1}},
		{head: pkg.MakeLinkedList(5, 3, 1, 2, 5, 1, 2), distances: []int{1, 3}},
		{head: pkg.MakeLinkedList(1, 3, 2, 2, 3, 2, 2, 2, 7), distances: []int{3, 3}},
	} {
		distances := nodesBetweenCriticalPoints(tc.head)
		assert.Equal(t, tc.distances, distances)
	}
}
