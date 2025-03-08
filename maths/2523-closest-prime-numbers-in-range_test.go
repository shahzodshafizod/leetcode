package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestClosestPrimes$
func TestClosestPrimes(t *testing.T) {
	for _, tc := range []struct {
		left  int
		right int
		ans   []int
	}{
		{left: 10, right: 19, ans: []int{11, 13}},
		{left: 4, right: 6, ans: []int{-1, -1}},
		{left: 1, right: 2, ans: []int{-1, -1}},
	} {
		ans := closestPrimes(tc.left, tc.right)
		assert.Equal(t, tc.ans, ans)
	}
}
