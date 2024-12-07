package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestFindNumbers$
func TestFindNumbers(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int
	}{
		{nums: []int{12, 345, 2, 6, 7896}, count: 2},
		{nums: []int{555, 901, 482, 1771}, count: 1},
		{nums: []int{437, 315, 322, 431, 686, 264, 442}, count: 0},
	} {
		count := findNumbers(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
