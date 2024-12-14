package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestPickGifts$
func TestPickGifts(t *testing.T) {
	for _, tc := range []struct {
		gifts    []int
		k        int
		remained int64
	}{
		{gifts: []int{1, 1, 1, 1}, k: 4, remained: 4},
		{gifts: []int{25, 64, 9, 4, 100}, k: 4, remained: 29},
		{gifts: []int{1, 2, 2, 2, 2, 2, 1}, k: 1000, remained: 7},
		{gifts: []int{11, 41, 66, 63, 14, 39, 70, 58, 12, 71}, k: 20, remained: 17},
		// {gifts: []int{64, 43, 32, 42, 9, 25, 73, 29, 56, 41, 27, 71, 62, 57, 67, 34, 8, 71, 2, 12, 52, 1, 31}, k: 51, remained: 33},
		// {gifts: []int{12, 34, 9, 29, 10, 39, 45, 56, 24, 8, 65, 48, 15, 5, 3, 25, 24, 16, 62, 27, 8, 3, 70, 55, 13, 60}, k: 3, remained: 591},
		// {gifts: []int{70, 51, 68, 36, 67, 31, 28, 54, 41, 74, 31, 38, 24, 25, 24, 5, 34, 61, 9, 12, 17, 20, 5, 11, 75}, k: 18, remained: 191},
		// {gifts: []int{8, 21, 21, 44, 68, 33, 16, 57, 23, 2, 61, 53, 73, 66, 40, 46, 50, 33, 20, 72, 2, 59, 11, 43, 6, 70, 13, 51, 26, 34, 46, 61, 73, 22, 27, 36, 18, 31, 62}, k: 24, remained: 396},
		// {gifts: []int{5, 52, 73, 54, 6, 22, 58, 9, 34, 21, 58, 68, 63, 72, 1, 5, 64, 42, 37, 46, 17, 40, 50, 54, 11, 1, 25, 43, 21, 31, 29, 58, 49, 73, 54, 40, 60, 7, 54, 25}, k: 37, remained: 214},
	} {
		remained := pickGifts(tc.gifts, tc.k)
		assert.Equal(t, tc.remained, remained)
	}
}
