package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestSortPeople$
func TestSortPeople(t *testing.T) {
	for _, tc := range []struct {
		names   []string
		heights []int
		sorted  []string
	}{
		{names: []string{"Mary", "John", "Emma"}, heights: []int{180, 165, 170}, sorted: []string{"Mary", "Emma", "John"}},
		{names: []string{"Alice", "Bob", "Bob"}, heights: []int{155, 185, 150}, sorted: []string{"Bob", "Alice", "Bob"}},
	} {
		sorted := sortPeople(tc.names, tc.heights)
		assert.Equal(t, tc.sorted, sorted)
	}
}
