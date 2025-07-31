package binarysearch

// https://leetcode.com/problems/maximum-candies-allocated-to-k-children/

func maximumCandies(candies []int, k int64) int {
	canDivide := func(target int64) bool {
		var count int64
		for _, candy := range candies {
			count += int64(candy) / target
			if count >= k {
				return true
			}
		}

		return false
	}

	var sum int64
	for _, candy := range candies {
		sum += int64(candy)
	}

	if sum < k {
		return 0
	}

	var low, high int64 = 1, sum / k

	var mid int64
	for low < high {
		mid = low + (high-low)/2 + 1
		if canDivide(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return int(low)
}
