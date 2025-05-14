package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestLengthAfterTransformations$
func TestLengthAfterTransformations(t *testing.T) {
	for _, tc := range []struct {
		s      string
		t      int
		nums   []int
		length int
	}{
		{s: "abcyy", t: 2, nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}, length: 7},
		{s: "azbk", t: 1, nums: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, length: 8},
		// {s: "zcxjrwgcqgmzgxyinfmnexquaopmzjllxlwtgxnpqvsgqbqvnguwymblhzkjkytdejhgewxtqwhpgmgrmvrazebvjikqplfcgnnosmxrcakvfvvjyvmkgzpglcpnphjttgmowbbfeqmqkbfszohhtrrjixnqctndm", t: 8942, nums: []int{21, 10, 18, 20, 22, 15, 19, 16, 1, 18, 17, 1, 22, 18, 3, 7, 13, 5, 21, 6, 10, 20, 22, 2, 25, 21}, length: 906321724},
	} {
		length := lengthAfterTransformations(tc.s, tc.t, tc.nums)
		assert.Equal(t, tc.length, length)
	}
}
