package bits

/*
^ is XOR (exclusive OR)
a ^ a = 0
a ^ 0 = a
*/

// https://leetcode.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/

// Approach 3: Nested Prefix XOR
// time: O(n ^ 2)
// space: O(1)
func countTriplets(arr []int) int {
	count := 0
	n := len(arr)

	var xor int

	for start := range n {
		xor = arr[start]
		for end := start + 1; end < n; end++ {
			xor ^= arr[end]
			if xor == 0 {
				count += end - start
			}
		}
	}

	return count
}

// // Approach 5: One Pass Prefix XOR
// // time: O(n)
// // space: O(2*n) = O(n)
// func countTriplets(arr []int) int {
// 	var prefix = 0
// 	var count = 0
// 	var prevXORcount = map[int]int{0: 1}
// 	var prevXORIndexSum = make(map[int]int)
// 	for i, n := 0, len(arr); i < n; i++ {
// 		prefix ^= arr[i]
// 		count += prevXORcount[prefix]*i - prevXORIndexSum[prefix]
// 		prevXORcount[prefix]++
// 		prevXORIndexSum[prefix] += i + 1
// 	}
// 	return count
// }

// // Approach 4: Two Pass Prefix XOR
// // time: O(n)
// // space: O(3*n) = O(n)
// func countTriplets(arr []int) int {
// 	var n = len(arr)
// 	var prefix = make([]int, n) // prefix XOR
// 	prefix[0] = arr[0]
// 	for i := 1; i < n; i++ {
// 		prefix[i] = prefix[i-1] ^ arr[i]
// 	}
// 	var count = 0
// 	var prevXORcount = map[int]int{0: 1}
// 	var prevXORIndexSum = make(map[int]int)
// 	for i, prefix := range prefix {
// 		count += prevXORcount[prefix]*i - prevXORIndexSum[prefix]
// 		prevXORcount[prefix]++
// 		prevXORIndexSum[prefix] += i + 1
// 	}
// 	return count
// }

// // Approach 2: Brute Force With Prefix
// // time: O(n ^ 3)
// // space: O(1)
// func countTriplets(arr []int) int {
// 	var n = len(arr)
// 	var count = 0
// 	var a, b int
// 	for i := 0; i < n-1; i++ {
// 		a = 0
// 		for j := i + 1; j < n; j++ {
// 			a ^= arr[j-1]
// 			b = 0
// 			for k := j; k < n; k++ {
// 				b ^= arr[j]
// 				if a == b {
// 					count++
// 				}
// 			}
// 		}
// 	}
// 	return count
// }

// // Approach 1: Brute Force
// // time: O(n ^ 4)
// // space: O(1)
// func countTriplets(arr []int) int {
// 	var n = len(arr)
// 	var count = 0
// 	var a, b int
// 	for i := 0; i < n-1; i++ {
// 		for j := i + 1; j < n; j++ {
// 			for k := j; k < n; k++ {
// 				a, b = 0, 0
// 				for idx := i; idx < j; idx++ {
// 					a ^= arr[idx]
// 				}
// 				for idx := j; idx <= k; idx++ {
// 					b ^= arr[idx]
// 				}
// 				if a == b {
// 					count++
// 				}
// 			}
// 		}
// 	}
// 	return count
// }
