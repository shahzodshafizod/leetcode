package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestKthSmallestPrimeFraction$
func TestKthSmallestPrimeFraction(t *testing.T) {
	for _, tc := range []struct {
		arr    []int
		k      int
		answer []int
	}{
		{arr: []int{1, 2, 3, 5}, k: 3, answer: []int{2, 5}},
		{arr: []int{1, 7}, k: 1, answer: []int{1, 7}},
	} {
		answer := kthSmallestPrimeFraction(tc.arr, tc.k)
		assert.Equal(t, tc.answer, answer)
	}
}
