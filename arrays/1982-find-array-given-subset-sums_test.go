package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestRecoverArray$
func TestRecoverArray(t *testing.T) {
	for _, tc := range []struct {
		n    int
		sums []int
		ans  []int
	}{
		{n: 3, sums: []int{-3, -2, -1, 0, 0, 1, 2, 3}, ans: []int{1, 2, -3}},
		{n: 2, sums: []int{0, 0, 0, 0}, ans: []int{0, 0}},
		{n: 4, sums: []int{0, 0, 5, 5, 4, -1, 4, 9, 9, -1, 4, 3, 4, 8, 3, 8}, ans: []int{0, -1, 4, 5}},
		{n: 3, sums: []int{0, 1, 3, 5, 4, 6, 8, 9}, ans: []int{1, 3, 5}},
		{n: 2, sums: []int{-605, 9, -596, 0}, ans: []int{9, -605}},
	} {
		ans := recoverArray(tc.n, tc.sums)
		assert.Equal(t, tc.ans, ans)
	}
}
