package binarysearch

// https://leetcode.com/problems/find-the-longest-valid-obstacle-course-at-each-position/

// Approach: Binary Search + LIS (Longest Increasing (Non-Decreasing) Subsequence)
// Time: O(n log n)
// Space: O(n)
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	lis := make([]int, n)
	size := 0
	ans := make([]int, 0, n)

	var left, right, mid int
	for _, obstacle := range obstacles {
		left, right = 0, size
		for left < right {
			mid = left + (right-left)/2
			if lis[mid] > obstacle {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if right == size {
			lis[size] = obstacle
			size++
		} else {
			lis[right] = obstacle
		}

		ans = append(ans, right+1)
	}

	return ans
}

// // Approach: Brute-Force
// // Time: O(nn)
// // Space: O(1)
// func longestObstacleCourseAtEachPosition(obstacles []int) []int {
// 	var n = len(obstacles)
// 	var ans = make([]int, n)
// 	var prev int
// 	for curr := range obstacles {
// 		prev = curr
// 		for j := curr - 1; j >= 0; j-- {
// 			if obstacles[j] <= obstacles[curr] && ans[j] > ans[prev] {
// 				prev = j
// 			}
// 		}
// 		if prev == curr {
// 			ans[curr] = 1
// 		} else {
// 			ans[curr] = ans[prev] + 1
// 		}
// 	}
// 	return ans
// }
