package prefixsums

// https://leetcode.com/problems/subarray-sum-equals-k/

// Approach #2: Prefix Sum + Hash Table
// Time: O(n)
// Space: O(n)
func subarraySum(nums []int, k int) int {
	counter := map[int]int{0: 1}

	count, presum := 0, 0
	for _, num := range nums {
		presum += num
		count += counter[presum-k]
		counter[presum]++
	}

	return count
}

// // Approach #1: Prefix Sum
// // Time: O(nn)
// // Space: O(1)
// func subarraySum(nums []int, k int) int {
// 	var count, presum = 0, 0
// 	var subsum int
// 	for end, num := range nums {
// 		presum += num
// 		subsum = presum
// 		for start := 0; start <= end; start++ {
// 			if subsum == k {
// 				count++
// 			}
// 			subsum -= nums[start]
// 		}
// 	}
// 	return count
// }
