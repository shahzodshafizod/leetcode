package heaps

import (
	"container/heap"
	"math"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/take-gifts-from-the-richest-pile/

func pickGifts(gifts []int, k int) int64 {
	var maxHeap = pkg.NewHeap(gifts, func(x, y int) bool { return x > y })
	heap.Init(maxHeap)
	for ; k > 0; k-- {
		gifts[0] = int(math.Sqrt(float64(gifts[0])))
		heap.Fix(maxHeap, 0)
	}
	var remained int64 = 0
	for _, gift := range gifts {
		remained += int64(gift)
	}
	return remained
}
