package slidingwindows

// https://leetcode.com/problems/count-number-of-nice-subarrays/

// Approach #3: Sliding Window (Three Pointers)
// time: O(2n) = O(n)
// space: O(1)
func numberOfSubarrays(nums []int, k int) int {
	subarrays := 0
	start, middle := 0, 0

	odds := 0
	for end := range nums {
		odds += nums[end] & 1
		for odds > k {
			odds -= nums[start] & 1
			start++
		}

		if odds == k {
			middle = start
			for nums[middle]&1 == 0 {
				middle++
			}

			subarrays += (middle - start) + 1
		}
	}

	return subarrays
}

// // Approach #2: Sliding Window (subarray sum at most k)
// // time: O(2n) = O(n)
// // space: O(1)
// func numberOfSubarrays(nums []int, k int) int {
// 	var atMost = func(k int) int {
// 		var subarrays = 0
// 		var start = 0
// 		var odds = 0
// 		for end := range nums {
// 			odds += nums[end] & 1
// 			for odds > k {
// 				odds -= nums[start] & 1
// 				start++
// 			}
// 			subarrays += end - start + 1
// 		}
// 		return subarrays
// 	}
// 	return atMost(k) - atMost(k-1)
// }

// // Approach #1: Hashing
// // time: O(n)
// // space: O(n)
// func numberOfSubarrays(nums []int, k int) int {
// 	var subarrays = 0
// 	var sum = 0
// 	var prefixSum = map[int]int{sum: 1}
// 	for _, num := range nums {
// 		sum += num % 2
// 		subarrays += prefixSum[sum-k]
// 		prefixSum[sum]++
// 	}
// 	return subarrays
// }
