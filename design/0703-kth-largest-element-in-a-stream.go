package design

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/kth-largest-element-in-a-stream/

type KthLargest struct {
	minHeap *pkg.Heap[int]
	len     int
}

func NewKthLargest(k int, nums []int) KthLargest {
	kth := KthLargest{len: k}
	compare := func(x, y int) bool { return x < y }
	if len(nums) <= k {
		kth.minHeap = pkg.NewHeap(nums, compare)
		nums = nil
	} else {
		kth.minHeap = pkg.NewHeap(nums[:k], compare)
		nums = nums[k:]
	}
	heap.Init(kth.minHeap)
	for _, num := range nums {
		if num > kth.minHeap.Peak() {
			heap.Pop(kth.minHeap)
			heap.Push(kth.minHeap, num)
		}
	}
	return kth
}

func (k *KthLargest) Add(val int) int {
	// take into account a case when heap_size is less than k
	if heap_size := k.minHeap.Len(); heap_size < k.len || val > k.minHeap.Peak() {
		if heap_size >= k.len {
			heap.Pop(k.minHeap)
		}
		heap.Push(k.minHeap, val)
	}
	return k.minHeap.Peak()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
