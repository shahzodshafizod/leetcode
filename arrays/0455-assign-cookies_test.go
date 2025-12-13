package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFindContentChildren$
func TestFindContentChildren(t *testing.T) {
	for _, tc := range []struct {
		g      []int
		s      []int
		output int
	}{
		{g: []int{1, 2, 3}, s: []int{1, 1}, output: 1},
		{g: []int{1, 2}, s: []int{1, 2, 3}, output: 2},
		{g: []int{1, 2, 3}, s: []int{3}, output: 1},
		{g: []int{10, 9, 8, 7}, s: []int{5, 6, 7, 8}, output: 2},
	} {
		output := findContentChildren(tc.g, tc.s)
		assert.Equal(t, tc.output, output)
	}
}
