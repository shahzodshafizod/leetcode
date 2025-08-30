package monotonic

import (
	"container/heap"
	"math"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/apply-operations-to-maximize-score/

// Approach: Monotonic Stack & Priority Queue
// Time: O(n*sqrt(max) + klogn)
// Space: O(4n) = O(n)
func maximumScore(nums []int, k int) int {
	n := len(nums)
	// step 1: calculate prime score
	scores := make([]int, n)
	maxHeap := make([][2]int, n)

	var limit int

	for idx, num := range nums {
		maxHeap[idx][0] = idx
		maxHeap[idx][1] = num

		limit = int(math.Sqrt(float64(num)))
		for f := 2; f <= limit; f++ {
			if num%f == 0 {
				scores[idx]++

				for num%f == 0 {
					num /= f
				}
			}
		}

		if num > 1 { // the number itself if prime
			scores[idx]++
		}
	}
	// step 2: find the previous greater or equal element
	left := make([]int, n)
	stack := make([]int, n)
	size := 0

	for idx := range n {
		for size > 0 && scores[stack[size-1]] < scores[idx] {
			size--
		}

		if size > 0 {
			left[idx] = stack[size-1]
		} else {
			left[idx] = -1
		}

		stack[size] = idx
		size++
	}
	// step 2: find the next greater element
	right := make([]int, n)

	size = 0
	for idx := n - 1; idx >= 0; idx-- {
		for size > 0 && scores[stack[size-1]] <= scores[idx] {
			size--
		}

		if size > 0 {
			right[idx] = stack[size-1]
		} else {
			right[idx] = n
		}

		stack[size] = idx
		size++
	}
	// step 3: Picking the highest prime score element
	pow := func(x, n int, mod int) int {
		res := 1

		for n > 0 {
			if n&1 == 1 {
				res = (res * x) % mod
			}

			x = (x * x) % mod
			n /= 2
		}

		return res
	}
	pq := pkg.NewHeap(maxHeap,
		func(x, y [2]int) bool { return x[1] > y[1] },
	)
	heap.Init(pq)

	score := 1

	var num, idx, count int

	const mod = int(1e9) + 7

	for k > 0 && pq.Len() > 0 {
		top, ok := heap.Pop(pq).([2]int)
		_ = ok
		idx, num = top[0], top[1]
		count = min((idx-left[idx])*(right[idx]-idx), k)
		k -= count
		score = (score * pow(num, count, mod)) % mod
	}

	return score
}
