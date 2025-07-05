package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestFindLucky$
func TestFindLucky(t *testing.T) {
	for _, tc := range []struct {
		arr   []int
		lucky int
	}{
		{arr: []int{2, 2, 3, 4}, lucky: 2},
		{arr: []int{1, 2, 2, 3, 3, 3}, lucky: 3},
		{arr: []int{2, 2, 2, 3, 3}, lucky: -1},
	} {
		lucky := findLucky(tc.arr)
		assert.Equal(t, tc.lucky, lucky)
	}
}
