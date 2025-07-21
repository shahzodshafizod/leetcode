package greedy

/*
Why we didn't use edges?
	Since it's a tree (nodes are connected), there is always a path between any two nodes.
	Suppose the path is A---B---C. We can only XOR both A and C by 2 operations:
		1. XOR A---B
		2. XOR B---C
	The node B is XORed twice and become to the original value due to XOR.
*/

// https://leetcode.com/problems/find-the-maximum-sum-of-node-values/

// Approach #6: Greedy (Finding local maxima and minima)
// Time: O(n)
// Space: O(1)
func maximumValueSum(nums []int, k int, _ [][]int) int64 {
	// 10^9 = 0b111011100110101100101000000000
	var sum, minDelta int64 = 0, 1 << 30
	var delta int64
	count := 0
	for _, num := range nums {
		sum += int64(num)
		delta = int64((num ^ k) - num)
		if delta > 0 {
			sum += int64(delta)
			count++
			minDelta = min(minDelta, delta)
		} else {
			minDelta = min(minDelta, -delta)
		}
	}
	if count&1 == 1 {
		sum -= minDelta
	}
	return sum
}

// // Approach #5: Greedy (Sorting based approach)
// // Time: O(nlogn)
// // Space: O(1)
// func maximumValueSum(nums []int, k int, edges [][]int) int64 {
// 	_ = edges
// 	var sum int64 = 0
// 	for idx := range nums {
// 		sum += int64(nums[idx])
// 		nums[idx] = (nums[idx] ^ k) - nums[idx]
// 	}
// 	sort.Ints(nums)
// 	for idx := len(nums) - 1; idx > 0 && nums[idx-1]+nums[idx] > 0; idx -= 2 {
// 		sum += int64(nums[idx] + nums[idx-1])
// 	}
// 	return sum
// }

// // Approach #4: Greedy (Sorting based approach)
// // Time: O(nlogn)
// // Space: O(n)
// func maximumValueSum(nums []int, k int, edges [][]int) int64 {
// 	_ = edges
// 	var n = len(nums)
// 	var delta = make([]int, n)
// 	var sum int64 = 0
// 	for idx := range delta {
// 		sum += int64(nums[idx])
// 		delta[idx] = (nums[idx] ^ k) - nums[idx]
// 	}
// 	sort.Ints(delta)
// 	for idx := n - 1; idx > 0 && delta[idx-1]+delta[idx] > 0; idx -= 2 {
// 		sum += int64(delta[idx] + delta[idx-1])
// 	}
// 	return sum
// }

// // Approach #3: Bottom-Up Dynamic Programming (Tabulation)
// // Time: O(n)
// // Space: O(n)
// func maximumValueSum(nums []int, k int, edges [][]int) int64 {
// 	_ = edges
// 	var n = len(nums)
// 	var dp = make([][2]int64, n+1)
// 	dp[n][0] = math.MinInt
// 	var xored, notxored int64
// 	for index := n - 1; index >= 0; index-- {
// 		for _, even := range [2]int{0, 1} {
// 			xored = dp[index+1][even^1] + int64(nums[index]^k)
// 			notxored = dp[index+1][even] + int64(nums[index])
// 			dp[index][even] = max(xored, notxored)
// 		}
// 	}
// 	return dp[0][1]
// }

// // Approach #2: Top-Down Dynamic Programming (Memoization)
// // Time: O(n)
// // Space: O(n)
// func maximumValueSum(nums []int, k int, edges [][]int) int64 {
// 	_ = edges
// 	var n = len(nums)
// 	var cache = make([][2]*int64, n)
// 	var dp func(idx int, even int) int64
// 	dp = func(idx int, even int) int64 {
// 		if idx == n {
// 			if even == 1 {
// 				return 0
// 			}
// 			return math.MinInt
// 		}
// 		if cache[idx][even] != nil {
// 			return *cache[idx][even]
// 		}
// 		var noXorDone = int64(nums[idx]) + dp(idx+1, even)
// 		var xorDone = int64(nums[idx]^k) + dp(idx+1, even^1)
// 		cache[idx][even] = new(int64)
// 		*cache[idx][even] = max(noXorDone, xorDone)
// 		return *cache[idx][even]
// 	}
// 	return dp(0, 1)
// }

// // Approach #1: Brute Force
// // Time: O(???)
// // Space: O(1)
// func maximumValueSum(nums []int, k int, edges [][]int) int64 {
// 	var u, v int
// 	var exit = false
// 	for !exit {
// 		exit = true
// 		for _, edge := range edges {
// 			u, v = edge[0], edge[1]
// 			if nums[u]+nums[v] < (nums[u]^k)+(nums[v]^k) {
// 				exit = false
// 				nums[u] ^= k
// 				nums[v] ^= k
// 			}
// 		}
// 	}
// 	var sum int64 = 0
// 	for _, num := range nums {
// 		sum += int64(num)
// 	}
// 	return sum
// }
