package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestFindEvenNumbers$
func TestFindEvenNumbers(t *testing.T) {
	for _, tc := range []struct {
		digits  []int
		uniques []int
	}{
		{digits: []int{2, 1, 3, 0}, uniques: []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320}},
		{digits: []int{2, 2, 8, 8, 2}, uniques: []int{222, 228, 282, 288, 822, 828, 882}},
		{digits: []int{3, 7, 5}, uniques: []int{}},
	} {
		uniques := findEvenNumbers(tc.digits)
		assert.ElementsMatch(t, tc.uniques, uniques)
	}
}
