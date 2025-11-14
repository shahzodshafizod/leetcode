package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestGetIntersectionNode$
func TestGetIntersectionNode(t *testing.T) {
	for _, tc := range []struct {
		listA []any
		listB []any
		skipA int
		skipB int
	}{
		{listA: []any{4, 1, 8, 4, 5}, listB: []any{5, 6, 1, 8, 4, 5}, skipA: 2, skipB: 3},
		{listA: []any{1, 9, 1, 2, 4}, listB: []any{3, 2, 4}, skipA: 3, skipB: 1},
		{listA: []any{2, 6, 4}, listB: []any{1, 5}, skipA: 3, skipB: 2},
	} {
		headA := pkg.MakeLinkedList(tc.listA...)
		headB := pkg.MakeLinkedList(tc.listB...)

		nodeA, nodeB := headA, headB
		for ; tc.skipA > 1; tc.skipA-- {
			nodeA = nodeA.Next
		}

		for ; tc.skipB > 1; tc.skipB-- {
			nodeB = nodeB.Next
		}

		nodeA.Next = nodeB.Next
		expected := nodeA.Next
		output := getIntersectionNode(headA, headB)
		assert.Equal(t, expected, output)
	}
}
