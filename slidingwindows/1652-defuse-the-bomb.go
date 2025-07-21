package slidingwindows

// https://leetcode.com/problems/defuse-the-bomb/

// Approach: Sliding Window
// Time: O(n+k), k=abs(k)
// Space: O(1)
func decrypt(code []int, k int) []int {
	n, start, end := len(code), 1, k
	if k < 0 {
		start, end = n+k, n-1
	}
	presum := 0
	for idx := start; idx <= end; idx++ {
		presum += code[idx]
	}
	decrypted := make([]int, n)
	for idx := 0; idx < n; idx++ {
		decrypted[idx] = presum
		presum = presum - code[start%n] + code[(end+1)%n]
		start, end = start+1, end+1
	}
	return decrypted
}

// // Approach: Brute-Force
// // Time: O(nk), k=abs(k)
// // Space: O(1)
// func decrypt(code []int, k int) []int {
// 	var n = len(code)
// 	var decrypted = make([]int, n)
// 	var step = 1
// 	if k > 0 {
// 		step = -1
// 	}
// 	for idx := 0; idx < n; idx++ {
// 		for j := idx + k; j != idx; j += step {
// 			decrypted[idx] += code[(j+n)%n]
// 		}
// 	}
// 	return decrypted
// }
