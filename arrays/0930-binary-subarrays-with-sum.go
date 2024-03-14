package arrays

// https://leetcode.com/problems/binary-subarrays-with-sum/

// time: O(n)
// space: O(1)
// (22 ms, 6.7 MB)
func numSubarraysWithSum(nums []int, goal int) int {
	var count, sum, zeros, start = 0, 0, 0, 0
	for end, num := range nums {
		sum += num
		for start < end && (nums[start] == 0 || sum > goal) {
			if nums[start] != 0 {
				zeros = 0
			} else {
				zeros++
			}
			sum -= nums[start]
			start++
		}
		if sum == goal {
			count += 1 + zeros
		}
	}
	return count
}

// // Sliding Window (31 ms, 6.6 MB)
// func numSubarraysWithSum(nums []int, goal int) int {
// 	var atMost = func(goal int) int {
// 		var start, sum, count = 0, 0, 0
// 		for end, num := range nums {
// 			sum += num
// 			for sum > goal && start <= end {
// 				sum -= nums[start]
// 				start++
// 			}
// 			count += end - start + 1
// 		}
// 		return count
// 	}
// 	return atMost(goal) - atMost(goal-1)
// }

// // Prefix Sum + Hash Table (32 ms, 7.8 MB)
// func numSubarraysWithSum(nums []int, goal int) int {
// 	var count = 0
// 	var sum = 0
// 	var freq = make(map[int]int)
// 	for _, num := range nums {
// 		sum += num
// 		if sum == goal {
// 			count++
// 		}
// 		if key := sum - goal; freq[key] > 0 {
// 			count += freq[key]
// 		}
// 		freq[sum]++
// 	}
// 	return count
// }

// // brute-force (1244 ms, 6.54 MB)
// func numSubarraysWithSum(nums []int, goal int) int {
// 	var sum int
// 	var count = 0
// 	for i, n := 0, len(nums); i < n; i++ {
// 		sum = 0
// 		for j := i; j < n && sum <= goal; j++ {
// 			sum += nums[j]
// 			if sum == goal {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
