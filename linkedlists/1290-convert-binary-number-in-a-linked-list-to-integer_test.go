package linkedlists

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./linkedlists/ -run ^TestGetDecimalValue$
func TestGetDecimalValue(t *testing.T) {
	for _, tc := range []struct {
		head  []any
		value int
	}{
		{head: []any{1, 0, 1}, value: 5},
		{head: []any{0}, value: 0},
	} {
		value := getDecimalValue(pkg.MakeLinkedList(tc.head...))
		assert.Equal(t, tc.value, value)
	}
}
