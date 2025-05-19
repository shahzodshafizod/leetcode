package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestTriangleType$
func TestTriangleType(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		typee string
	}{
		{nums: []int{3, 3, 3}, typee: "equilateral"},
		{nums: []int{3, 4, 5}, typee: "scalene"},
		{nums: []int{1, 1, 1}, typee: "equilateral"},
		{nums: []int{3, 1, 2}, typee: "none"},
		{nums: []int{2, 5, 2}, typee: "none"},
		{nums: []int{3, 3, 1}, typee: "isosceles"},
		{nums: []int{1, 1, 2}, typee: "none"},
		{nums: []int{2, 3, 4}, typee: "scalene"},
		{nums: []int{8, 4, 2}, typee: "none"},
		{nums: []int{8, 4, 4}, typee: "none"},
	} {
		typee := triangleType(tc.nums)
		assert.Equal(t, tc.typee, typee)
	}
}
