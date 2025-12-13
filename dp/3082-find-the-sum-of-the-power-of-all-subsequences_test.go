package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestSumOfPower$
func TestSumOfPower(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		output int
	}{
		{nums: []int{1, 2, 3}, k: 3, output: 6},
		{nums: []int{2, 3, 3}, k: 5, output: 4},
		{nums: []int{1, 2, 3}, k: 7, output: 0},
		{nums: []int{15, 7, 9, 8, 22, 15, 16, 21, 9, 12, 8, 24, 12, 6, 12, 18, 4, 25, 14, 17, 1, 15, 1, 5, 18}, k: 15, output: 148897792},
		{nums: []int{14, 19, 7, 24, 21, 4, 21, 5, 20, 2, 3, 17, 8, 3, 16, 3, 24, 7, 10, 24, 9, 11, 24, 11, 19}, k: 24, output: 487587840},
	} {
		output := sumOfPower(tc.nums, tc.k)
		assert.Equal(t, tc.output, output)
	}
}
