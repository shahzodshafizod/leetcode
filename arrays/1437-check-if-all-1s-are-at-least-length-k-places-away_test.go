package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestKLengthApart$
func TestKLengthApart(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		output bool
	}{
		{nums: []int{1, 0, 0, 0, 1, 0, 0, 1}, k: 2, output: true},
		{nums: []int{1, 0, 0, 1, 0, 1}, k: 2, output: false},
		{nums: []int{1, 1, 1, 1, 1}, k: 0, output: true},
		{nums: []int{0, 1, 0, 1}, k: 1, output: true},
		{nums: []int{0}, k: 1, output: true},
		{nums: []int{1}, k: 0, output: true},
		{nums: []int{1, 0, 1}, k: 1, output: true},
		{nums: []int{1, 0, 1}, k: 2, output: false},
	} {
		output := kLengthApart(tc.nums, tc.k)
		assert.Equal(t, tc.output, output)
	}
}
