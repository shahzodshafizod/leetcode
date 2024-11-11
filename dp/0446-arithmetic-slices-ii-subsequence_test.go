package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumberOfArithmeticSlicesSubsequence$
func TestNumberOfArithmeticSlicesSubsequence(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int
	}{
		{nums: []int{-104}, count: 0},
		{nums: []int{7, 7, 7, 7, 7}, count: 16},
		{nums: []int{2, 4, 6, 8, 10}, count: 7},
		{nums: []int{0, 2000000000, -294967296}, count: 0},
		{nums: []int{211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211}, count: 16776915},
		{nums: []int{-454, -908, -1362, -454, -908, -1362, -454, -908, -1362, -454, -908, -1362, -454, -908, -1362, -454, -908, -1362, -454, -908, -1362, -454, -908, -1362, -454, -908, -1362, -454, -908, -1362}, count: 3244},
		{nums: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912, 1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683}, count: 1},
		{nums: []int{1073741824, 536870912, 268435456, 134217728, 67108864, 33554432, 16777216, 8388608, 4194304, 2097152, 1048576, 524288, 262144, 131072, 65536, 32768, 16384, 8192, 4096, 2048, 1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1073740824, 536870412, 268435206, 134217603, 67108801, 33554400, 16777200, 8388600, 4194300, 2097150, 1048575, 524287, 262143, 131071, 65535, 32767, 16383, 8191, 4095, 2047, 1023, 511, 255, 127, 63, 31, 15, 7, 3, 1}, count: 0},
	} {
		count := numberOfArithmeticSlicesSubsequence(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}