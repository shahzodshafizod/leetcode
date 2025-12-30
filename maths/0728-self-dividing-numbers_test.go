package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestSelfDividingNumbers$
func TestSelfDividingNumbers(t *testing.T) {
	for _, tc := range []struct {
		left   int
		right  int
		output []int
	}{
		{left: 1, right: 22, output: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{left: 47, right: 85, output: []int{48, 55, 66, 77}},
		{left: 1, right: 1, output: []int{1}},
		{left: 10, right: 20, output: []int{11, 12, 15}},
		{left: 21, right: 30, output: []int{22, 24}},
	} {
		output := selfDividingNumbers(tc.left, tc.right)
		assert.Equal(t, tc.output, output)
	}
}
