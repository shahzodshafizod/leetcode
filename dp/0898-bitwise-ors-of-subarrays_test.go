package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestSubarrayBitwiseORs$
func TestSubarrayBitwiseORs(t *testing.T) {
	for _, tc := range []struct {
		arr []int
		cnt int
	}{
		{arr: []int{0}, cnt: 1},
		{arr: []int{1, 1, 2}, cnt: 3},
		{arr: []int{1, 2, 4}, cnt: 6},
	} {
		cnt := subarrayBitwiseORs(tc.arr)
		assert.Equal(t, tc.cnt, cnt)
	}
}
