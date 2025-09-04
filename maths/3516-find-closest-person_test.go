package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestFindClosest$
func TestFindClosest(t *testing.T) {
	for _, tc := range []struct {
		x int
		y int
		z int
		num int
	}{
		{x: 2, y: 7, z: 4, num: 1},
		{x: 2, y: 5, z: 6, num: 2},
		{x: 1, y: 5, z: 3, num: 0},
	} {
		num := findClosest(tc.x, tc.y, tc.z)
		assert.Equal(t, tc.num, num)
	}
}
