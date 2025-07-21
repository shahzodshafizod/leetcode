package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestFindInMountainArray$
func TestFindInMountainArray(t *testing.T) {
	for _, tc := range []struct {
		target      int
		mountainArr []int
		index       int
	}{
		{mountainArr: []int{6, 7, 2}, target: 2, index: 2},
		{mountainArr: []int{0, 1, 2, 4, 2, 1}, target: 3, index: -1},
		{mountainArr: []int{1, 2, 3, 4, 5, 3, 1}, target: 3, index: 2},
		{mountainArr: []int{7, 17, 24, 28, 30, 33, 26, 20, 12, 2}, target: 2, index: 9},
		// {mountainArr: []int{11, 21, 29, 38, 39, 40, 47, 49, 50, 54, 57, 60, 61, 69, 78, 81, 83, 92, 102, 104, 109, 114, 123, 132, 135, 143, 144, 153, 163, 166, 167, 175, 173, 165, 155, 145, 136, 126, 118, 110, 104, 103, 100, 92, 87, 77, 71, 65, 62, 59, 57, 52, 49, 41, 40, 34, 30, 28, 24, 15, 7, 0}, target: 92, index: 17},
	} {
		mountainArr := &MountainArray{tc.mountainArr}
		index := findInMountainArray(tc.target, mountainArr)
		assert.Equal(t, tc.index, index)
	}
}
