package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestKthDistinct$
func TestKthDistinct(t *testing.T) {
	for _, tc := range []struct {
		arr      []string
		k        int
		distinct string
	}{
		{arr: []string{"d", "b", "c", "b", "c", "a"}, k: 2, distinct: "a"},
		{arr: []string{"aaa", "aa", "a"}, k: 1, distinct: "aaa"},
		{arr: []string{"a", "b", "a"}, k: 3, distinct: ""},
	} {
		distinct := kthDistinct(tc.arr, tc.k)
		assert.Equal(t, tc.distinct, distinct)
	}
}
