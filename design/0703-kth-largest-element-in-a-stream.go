package design

// https://leetcode.com/problems/kth-largest-element-in-a-stream/

type KthLargest struct {
	minHeap PriorityQueue[int]
	len     int
}

func NewKthLargest(k int, nums []int) KthLargest {
	var kth = KthLargest{len: k}
	var compare = func(x, y int) bool { return x > y }
	if len(nums) <= k {
		kth.minHeap = NewPriorityQueue(nums, compare)
		nums = nil
	} else {
		kth.minHeap = NewPriorityQueue(nums[:k], compare)
		nums = nums[k:]
	}
	kth.minHeap.Heapify()
	for _, num := range nums {
		if num > kth.minHeap.Peek() {
			kth.minHeap.Pop()
			kth.minHeap.Push(num)
		}
	}
	return kth
}

func (k *KthLargest) Add(val int) int {
	// take into account a case when heap_size is less than k
	if heap_size := k.minHeap.Len(); heap_size < k.len || val > k.minHeap.Peek() {
		if heap_size >= k.len {
			k.minHeap.Pop()
		}
		k.minHeap.Push(val)
	}
	return k.minHeap.Peek()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
