package slidingwindows

// https://leetcode.com/problems/maximum-fruits-harvested-after-at-most-k-steps/

// func maxTotalFruits(fruits [][]int, startPos int, k int) int {
// 	steps := func(left int, right int) int {
// 		if fruits[right][0] <= startPos {
// 			return startPos - fruits[left][0]
// 		}

// 		if startPos <= fruits[left][0] {
// 			return fruits[right][0] - startPos
// 		}

// 		return min(
// 			fruits[right][0]-startPos,
// 			startPos-fruits[left][0],
// 		) + fruits[right][0] - fruits[left][0]
// 	}

// 	res, total := 0, 0
// 	for start, end, n := 0, 0, len(fruits); end < n; end++ {
// 		total += fruits[end][1]
// 		for start <= end && steps(start, end) > k {
// 			total -= fruits[start][1]
// 			start++
// 		}

// 		res = max(res, total)
// 	}

// 	return res
// }

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	// to find left, could be used binary search approach
	left, n := 0, len(fruits)
	for left < n && fruits[left][0] < startPos-k {
		left++
	}

	total, maxTotal := 0, 0
	for right := left; right < n && fruits[right][0] <= startPos+k; right++ {
		total += fruits[right][1]
		for min(
			startPos-2*fruits[left][0]+fruits[right][0],
			2*fruits[right][0]-fruits[left][0]-startPos,
		) > k {
			total -= fruits[left][1]
			left++
		}

		maxTotal = max(maxTotal, total)
	}

	return maxTotal
}

// // Time: O(N)
// // Space: O(N)
// func maxTotalFruits(fruits [][]int, startPos int, k int) int {
// 	var (
// 		n      = len(fruits)
// 		maxPos = startPos + k + 1
// 		presum = make([]int, maxPos+1)
// 		idx    = 0
// 	)

// 	for pos := 1; pos <= maxPos; pos++ {
// 		if idx < n && pos-1 == fruits[idx][0] {
// 			presum[pos] = fruits[idx][1]
// 			idx++
// 		}

// 		presum[pos] += presum[pos-1]
// 	}

// 	harvest := 0
// 	// left-then-right
// 	var left, right int
// 	for left = startPos - k; left <= startPos; left++ {
// 		right = max(startPos, startPos+k-2*(startPos-left))
// 		harvest = max(harvest, presum[right+1]-presum[max(0, left)])
// 	}
// 	// right-then-left
// 	for right = startPos + k; right >= startPos; right-- {
// 		left = min(startPos, startPos-(k-2*(right-startPos)))
// 		harvest = max(harvest, presum[right+1]-presum[max(0, left)])
// 	}

// 	return harvest
// }
