package dp

// https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array/

// Approach: Dynamic Programming (Bottom-Up)
// Time: O(nlogn)
// Space: O(n)
func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	getLIS := func(start int, end int, step int) []int {
		lis := make([]int, n)
		sequence := make([]int, 0)
		slen := 0

		var left, right, mid int

		for idx := start; idx != end; idx += step {
			if slen == 0 || nums[idx] > sequence[slen-1] {
				sequence = append(sequence, nums[idx])
				slen++
				lis[idx] = slen

				continue
			}

			left, right = 0, slen
			for left < right {
				mid = left + (right-left)/2
				if sequence[mid] < nums[idx] {
					left = mid + 1
				} else {
					right = mid
				}
			}

			sequence[left] = nums[idx]
			lis[idx] = left + 1
		}

		return lis
	}
	// longest increasing subsequence
	lis := getLIS(0, n, 1)
	// longest decreasing subsequence
	lds := getLIS(n-1, -1, -1)

	var length int

	mountain := 0
	// minimum to remove = maximum to keep
	for peek, n := 0, len(nums); peek < n; peek++ { // O(n)
		length = lis[peek] + lds[peek] - 1
		if lis[peek] > 1 && lds[peek] > 1 && length >= 3 {
			mountain = max(mountain, length)
		}
	}

	return n - mountain
}

// // Approach: Dynamic Programming (Bottom-Up)
// // Time: O(nn)
// // Space: O(n)
// func minimumMountainRemovals(nums []int) int {
// 	var n = len(nums)
// 	// longest increasing subsequence
// 	var lis = make([]int, n)
// 	for curr := 0; curr < n; curr++ { // O(nn)
// 		lis[curr] = 1
// 		for prev := curr - 1; prev >= 0; prev-- {
// 			if nums[prev] < nums[curr] {
// 				lis[curr] = max(lis[curr], 1+lis[prev])
// 			}
// 		}
// 	}
// 	// longest decreasing subsequence
// 	var lds = make([]int, n)
// 	for curr := n - 1; curr >= 0; curr-- { // O(nn)
// 		lds[curr] = 1
// 		for next := curr + 1; next < n; next++ {
// 			if nums[curr] > nums[next] {
// 				lds[curr] = max(lds[curr], 1+lds[next])
// 			}
// 		}
// 	}
// 	var length int
// 	var mountain = 0
// 	// minimum to remove = maximum to keep
// 	for peek, n := 0, len(nums); peek < n; peek++ { // O(n)
// 		length = lis[peek] + lds[peek] - 1
// 		if lis[peek] > 1 && lds[peek] > 1 && length >= 3 {
// 			mountain = max(mountain, length)
// 		}
// 	}
// 	return n - mountain
// }
