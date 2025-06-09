package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestLexicalOrder$
func TestLexicalOrder(t *testing.T) {
	for _, tc := range []struct {
		n    int
		nums []int
	}{
		{n: 13, nums: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
		{n: 2, nums: []int{1, 2}},
	} {
		nums := lexicalOrder(tc.n)
		assert.Equal(t, tc.nums, nums)
	}
}
