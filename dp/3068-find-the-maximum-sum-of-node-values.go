package dp

import "math"

/*
Why we didn't use edges?
	Since it's a tree (nodes are connected), there is always a path between any two nodes.
	Suppose the path is A---B---C. We can only XOR both A and C by 2 operations:
		1. XOR A---B
		2. XOR B---C
	The node B is XORed twice and become to the original value due to XOR.
*/

// https://leetcode.com/problems/find-the-maximum-sum-of-node-values/

// time: O(n)
// space: O(1)
func maximumValueSum(nums []int, k int, _ [][]int) int64 {
	var sum int64 = 0
	var delta int
	var posmin = math.MaxInt
	var negmax = math.MinInt
	var count = 0
	for _, num := range nums {
		sum += int64(num)
		delta = (num ^ k) - num
		if delta > 0 {
			posmin = min(posmin, delta)
			sum += int64(delta)
			count++
		} else {
			negmax = max(negmax, delta)
		}
	}
	if count%2 != 0 {
		sum = max(sum-int64(posmin), sum+int64(negmax))
	}
	return sum
}

// // time: O(n log n)
// // space: O(1)
// func maximumValueSum(nums []int, k int, edges [][]int) int64 {
// 	var sum int64 = 0
// 	for idx := range nums {
// 		sum += int64(nums[idx])
// 		nums[idx] = (nums[idx] ^ k) - nums[idx]
// 	}
// 	sort.Ints(nums)
// 	var pairSum int64
// 	for idx := len(nums) - 1; idx > 0; idx -= 2 {
// 		pairSum = int64(nums[idx] + nums[idx-1])
// 		if pairSum > 0 {
// 			sum += pairSum
// 		}
// 	}
// 	return sum
// }

// // time: O(n log n)
// // space: O(n)
// func maximumValueSum(nums []int, k int, edges [][]int) int64 {
// 	var n = len(nums)
// 	var delta = make([]int, n)
// 	var sum int64 = 0
// 	for idx := range delta {
// 		sum += int64(nums[idx])
// 		delta[idx] = (nums[idx] ^ k) - nums[idx]
// 	}
// 	sort.Ints(delta)
// 	var pairSum int64
// 	for idx := n - 1; idx > 0; idx -= 2 {
// 		pairSum = int64(delta[idx] + delta[idx-1])
// 		if pairSum > 0 {
// 			sum += pairSum
// 		}
// 	}
// 	return sum
// }

// // Tabulation
// // time: O(n)
// // space: O(2n) = O(n)
// func maximumValueSum(nums []int, k int, edges [][]int) int64 {
// 	var n = len(nums)
// 	var dp = make([][2]int64, n+1)
// 	dp[n][0] = math.MinInt
// 	var xored, notxored int64
// 	for index := n - 1; index >= 0; index-- {
// 		for isEven := 0; isEven < 2; isEven++ {
// 			xored = dp[index+1][isEven^1] + int64(nums[index]^k)
// 			notxored = dp[index+1][isEven] + int64(nums[index])
// 			dp[index][isEven] = max(xored, notxored)
// 		}
// 	}
// 	return dp[0][1]
// }

// // Memoization
// // time: O(n)
// // space: O(2*n) = O(n)
// func maximumValueSum(nums []int, k int, edges [][]int) int64 {
// 	var n = len(nums)
// 	var cache = make([][2]int64, n)
// 	for idx := range cache {
// 		cache[idx][0] = -1
// 		cache[idx][1] = -1
// 	}
// 	var maxSumOfNodes func(index int, isEven int) int64
// 	maxSumOfNodes = func(index int, isEven int) int64 {
// 		if index == n {
// 			if isEven == 1 {
// 				return 0
// 			}
// 			return math.MinInt
// 		}
// 		if cache[index][isEven] != -1 {
// 			return cache[index][isEven]
// 		}
// 		var noXorDone = int64(nums[index]) + maxSumOfNodes(index+1, isEven)
// 		var xorDone = int64(nums[index]^k) + maxSumOfNodes(index+1, isEven^1)
// 		cache[index][isEven] = max(noXorDone, xorDone)
// 		return cache[index][isEven]
// 	}
// 	return maxSumOfNodes(0, 0)
// }

// // brute-force
// // time: O(n*k) ???
// // space: O(1)
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
