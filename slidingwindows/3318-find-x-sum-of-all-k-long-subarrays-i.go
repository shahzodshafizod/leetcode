package slidingwindows

import (
	"sort"
)

// https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-i/

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	answer := make([]int, n-k+1)
	counter := make(map[int]int)

	var num, cnt int

	for i := range nums {
		counter[nums[i]]++
		if i >= k-1 {
			var xsum [][2]int

			for num, cnt = range counter {
				if cnt > 0 {
					xsum = append(xsum, [2]int{cnt, num})
				}
			}

			sort.Slice(xsum, func(i, j int) bool {
				if xsum[i][0] == xsum[j][0] {
					return xsum[i][1] > xsum[j][1]
				}

				return xsum[i][0] > xsum[j][0]
			})

			for j := min(x, len(xsum)) - 1; j >= 0; j-- {
				cnt, num = xsum[j][0], xsum[j][1]
				answer[i-k+1] += num * cnt
			}

			counter[nums[i-k+1]]--
		}
	}

	return answer
}
