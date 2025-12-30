package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestIsSelfCrossing$
func TestIsSelfCrossing(t *testing.T) {
	for _, tc := range []struct {
		distance []int
		output   bool
	}{
		{distance: []int{2, 1, 1, 2}, output: true},
		{distance: []int{1, 2, 3, 4}, output: false},
		{distance: []int{1, 1, 1, 2, 1}, output: true},
		{distance: []int{1, 1, 2, 1, 1}, output: true},
		{distance: []int{3, 3, 4, 2, 2}, output: false},
		{distance: []int{1, 1, 2, 2, 3, 3, 4, 4, 10, 4, 4, 3, 3, 2, 2, 1, 1}, output: false},
		{distance: []int{1, 1, 2, 2, 3, 3, 4, 4, 10, 4, 4, 3, 3, 2, 2, 1, 1, 2}, output: true},
		{distance: []int{3, 3, 3, 2, 1, 1}, output: false},
	} {
		output := isSelfCrossing(tc.distance)
		assert.Equal(t, tc.output, output)
	}
}
