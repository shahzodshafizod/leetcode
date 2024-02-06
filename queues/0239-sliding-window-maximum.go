package queues

// https://leetcode.com/problems/sliding-window-maximum/

// func maxSlidingWindow(nums []int, k int) []int {
// 	var len = len(nums)
// 	if len == 0 || k == 0 {
// 		return []int{}
// 	}
// 	var list = list.New()
// 	var max = make([]int, 0)
// 	for i, num := range nums {
// 		// Hint 2: The queue size need not be the same as the window’s size.
// 		// Hint 3: Remove redundant elements and the queue should store only elements that need to be considered.
// 		for list.Len() > 0 && nums[list.Back().Value.(int)] <= num {
// 			list.Remove(list.Back())
// 		}
// 		list.PushBack(i)
// 		if i >= k-1 {
// 			max = append(max, nums[list.Front().Value.(int)])
// 			if list.Len() > 0 && list.Front().Value.(int) == i-k+1 {
// 				list.Remove(list.Front())
// 			}
// 		}
// 	}
// 	return max
// }

type deque struct { // double-ended queue
	index int
	next  *deque
	prev  *deque
}

type window struct {
	first *deque
	last  *deque
}

func (w *window) push(nums []int, index int) {
	// Hint 2: The queue size need not be the same as the window’s size.
	// Hint 3: Remove redundant elements and the queue should store only elements that need to be considered.
	for w.last != nil && nums[w.last.index] <= nums[index] {
		if w.last.prev != nil {
			w.last.prev.next = nil
		}
		w.last = w.last.prev
	}
	if w.last == nil {
		w.first = nil
	}
	var newNode = &deque{index: index, prev: w.last}
	if w.first == nil {
		w.first, w.last = newNode, newNode
	} else {
		w.last.next = newNode
		w.last = newNode
	}
}

func (w *window) pop(index int) {
	if w.first == nil || w.first.index != index {
		return
	}
	if w.first == w.last {
		w.first, w.last = nil, nil
	} else {
		w.first.next.prev = nil
		w.first = w.first.next
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	var len = len(nums)
	if len == 0 || k == 0 {
		return []int{}
	}
	if len == 1 {
		return nums
	}
	var window = &window{}
	var maxes = make([]int, 0, len-k+1)
	for i := range nums {
		window.push(nums, i)
		if i >= k-1 {
			maxes = append(maxes, nums[window.first.index])
			window.pop(i - k + 1)
		}
	}
	return maxes
}
