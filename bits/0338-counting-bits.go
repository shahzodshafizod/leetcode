package bits

// https://leetcode.com/problems/counting-bits/

func countBits(n int) []int {
	var bits = make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 { // power of two, has the only bit set
			bits[i] = 1
		} else if i&1 == 1 { // if odd
			bits[i] = bits[i-1] + 1
		} else {
			bits[i] = bits[i/2]
		}
	}
	return bits
}

// func countBits(n int) []int {
// 	var bits = make([]int, n+1)
// 	for i := 1; i <= n; i++ {
// 		bits[i] = bits[i/2] + i%2 // OR bits[i/2] + (i&1)
// 	}
// 	return bits
// }

// func countBits(n int) []int {
// 	var ones = make([]int, n+1)
// 	for i := 1; i <= n; i++ {
// 		ones[i] = bits.OnesCount(uint(i)) // math/bits
// 	}
// 	return ones
// }
