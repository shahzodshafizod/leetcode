package slidingwindows

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-ii/

// Approach #2: Two Heaps
// Time: O(n log k)
// Space: O(k)
func findXSumII(nums []int, k int, x int) []int64 {
	type elem struct {
		cnt int
		num int
		vrs int
	}

	fm := make(map[int]int) // frequency map: num -> frequency

	lsCmp := func(a, b *elem) bool {
		return a.cnt < b.cnt || a.cnt == b.cnt && a.num < b.num
	}

	// top is min heap - smallest of the x largest elements is at top
	top := pkg.NewHeap([]*elem{}, lsCmp)
	// low is max heap - largest of the remaining elements is at top
	low := pkg.NewHeap([]*elem{}, func(x, y *elem) bool { return !lsCmp(x, y) })

	versions := make(map[int]int)

	isInTop := make(map[int]bool)

	var cntInTop int

	var curr int64 // current sum of top x elements

	change := func(num int, count int) {
		// Remove previous state if it exists
		if prevFreq, exists := fm[num]; exists && prevFreq > 0 {
			if isInTop[num] {
				curr -= int64(prevFreq * num)
				isInTop[num] = false
				cntInTop--
			}
		}

		fm[num] += count
		versions[num]++

		// Add new state if frequency > 0
		if fm[num] > 0 {
			heap.Push(low, &elem{
				cnt: fm[num],
				num: num,
				vrs: versions[num],
			})
		}

		// Rebalance: move from low to top if top doesn't have enough valid elements
		for low.Len() > 0 && cntInTop < x {
			e, ok := heap.Pop(low).(*elem)
			_ = ok

			if e.vrs == versions[e.num] {
				heap.Push(top, e)
				curr += int64(e.cnt * e.num)

				isInTop[e.num] = true
				cntInTop++
			}
		}

		// Rebalance: swap if low has greater elements than top's minimum
		for low.Len() > 0 && top.Len() > 0 {
			lowMax := low.Peak()
			topMin := top.Peak()

			if lowMax.vrs < versions[lowMax.num] {
				heap.Pop(low)

				continue
			}

			if topMin.vrs < versions[topMin.num] {
				heap.Pop(top)

				continue
			}

			// check if no need for swapping
			if lsCmp(lowMax, topMin) {
				break
			}

			// Swap elements
			heap.Pop(low)
			heap.Pop(top)

			heap.Push(top, lowMax)
			heap.Push(low, topMin)

			// Update curr and tracking
			curr = curr - int64(topMin.cnt*topMin.num) + int64(lowMax.cnt*lowMax.num)
			isInTop[topMin.num] = false
			isInTop[lowMax.num] = true
		}
	}

	n := len(nums)
	answer := make([]int64, n-k+1)

	for i := range n {
		change(nums[i], 1)

		if i >= k-1 {
			answer[i-k+1] = curr
			change(nums[i-k+1], -1)
		}
	}

	return answer
}

// // Brute-force solution
// func findXSumII(nums []int, k int, x int) []int64 {
// 	n := len(nums)
// 	answer := make([]int64, 0, n-k+1)

// 	for i := 0; i <= n-k; i++ {
// 		// Count frequencies in current window
// 		freq := make(map[int]int)
// 		for j := i; j < i+k; j++ {
// 			freq[nums[j]]++
// 		}

// 		sorted := make([][2]int, 0, len(freq))
// 		for num, cnt := range freq {
// 			sorted = append(sorted, [2]int{cnt, num})
// 		}

// 		sort.Slice(sorted, func(i, j int) bool {
// 			return sorted[i][0] > sorted[j][0] ||
// 				sorted[i][0] == sorted[j][0] && sorted[i][1] > sorted[j][1]
// 		})

// 		var xSum int64

// 		for j := range min(x, len(sorted)) {
// 			xSum += int64(sorted[j][0]) * int64(sorted[j][1])
// 		}

// 		answer = append(answer, xSum)
// 	}

// 	return answer
// }
