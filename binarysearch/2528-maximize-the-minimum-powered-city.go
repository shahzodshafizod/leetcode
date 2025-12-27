package binarysearch

// https://leetcode.com/problems/maximize-the-minimum-powered-city/

// Approach: Binary Search
// Time: O(n log (s+k)), s=sum(stations)
// Space: O(n)
func maxPower(stations []int, r int, k int) int64 {
	var (
		n           = len(stations)
		left, right int
		mini, total int64 = 1e5, 0
		power       int64
	)

	line := make([]int64, n+1)

	for i := range stations {
		power = int64(stations[i])
		left = max(i-r, 0)
		right = min(i+r, n-1)
		line[left] += power
		line[right+1] -= power
		mini = min(mini, power)
		total += power
	}

	maximize := func(target int64) bool {
		presum, quota := int64(0), int64(k)

		diff := make([]int64, n+1)
		copy(diff, line)

		for i := range n {
			presum += diff[i]
			if presum+quota < target {
				return false
			}

			if presum < target {
				delta := target - presum
				quota -= delta
				presum += delta
				right = min(i+2*r, n-1)
				diff[right+1] -= delta
			}
		}

		return true
	}

	var mid int64

	low, high := mini, total+int64(k)

	ans := low
	for low <= high {
		mid = low + (high-low)/2
		if maximize(mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}
