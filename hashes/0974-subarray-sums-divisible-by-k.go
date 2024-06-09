package hashes

// https://leetcode.com/problems/subarray-sums-divisible-by-k/

// time: O(n)
// space: O(k)
func subarraysDivByK(nums []int, k int) int {
	var subarrays = 0
	var modGroups = make([]int, k)
	modGroups[0] = 1
	var prefixMod = 0
	for _, num := range nums {
		// take modulo twice to avoid negative remainders.
		prefixMod = (prefixMod + num%k + k) % k
		subarrays += modGroups[prefixMod]
		modGroups[prefixMod]++
	}
	return subarrays
}

// // time: O(n)
// // space: O(k)
// func subarraysDivByK(nums []int, k int) int {
// 	var subarrays = 0
// 	var modGroups = make([]int, k)
// 	modGroups[0] = 1
// 	var prefixSum = 0
// 	var mod int
// 	for _, num := range nums {
// 		prefixSum += num
// 		mod = prefixSum % k
// 		if mod < 0 {
// 			mod += k
// 		}
// 		subarrays += modGroups[mod]
// 		modGroups[mod]++
// 	}
// 	return subarrays
// }

// // time: O(n ^ 2)
// // space: O(1)
// func subarraysDivByK(nums []int, k int) int {
// 	var count = 0
// 	var sum int
// 	for i, n := 0, len(nums); i < n; i++ {
// 		sum = 0
// 		for j := i; j < n; j++ {
// 			sum += nums[j]
// 			if sum%k == 0 {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
