package maths

// https://leetcode.com/problems/count-array-pairs-divisible-by-k/

// Time: O(n*sqrt(K))
// Space: O(k)
func countPairs(nums []int, k int) int64 {
	gcd := func(a int, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	counts := make(map[int]int64)
	var count int64 = 0
	for _, num := range nums {
		curr := gcd(num, k)
		for prev, cnt := range counts {
			if curr*prev%k == 0 {
				count += cnt
			}
		}
		counts[curr]++
	}
	return count
}

// // Time: O(n*sqrt(k))
// // Space: O(k)
// func countPairs(nums []int, k int) int64 {
// 	var gcd = func(a int, b int) int {
// 		for b != 0 {
// 			a, b = b, a%b
// 		}
// 		return a
// 	}
// 	var divisors = make([]int, 0)
// 	for num := 1; num <= k; num++ {
// 		if k%num == 0 {
// 			divisors = append(divisors, num)
// 		}
// 	}
// 	var counter = make(map[int]int64)
// 	var count int64 = 0
// 	for _, num := range nums {
// 		rem := k / gcd(num, k)
// 		count += counter[rem]
// 		for _, divisor := range divisors {
// 			if num%divisor == 0 {
// 				counter[divisor]++
// 			}
// 		}
// 	}
// 	return count
// }

// // Approach: Brute-Force
// // Time: O(N^2)
// // Space: O(1)
// func countPairs(nums []int, k int) int64 {
// 	var count int64 = 0
// 	for i, n := 0, len(nums); i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			if nums[i]*nums[j]%k == 0 {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
