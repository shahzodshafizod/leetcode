package bits

/*
XOR: ^
0 ^ 0 = 0
1 ^ 1 = 0
0 ^ 1 = 1
1 ^ 0 = 1
*/

// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-xor-equal-to-k/

func minOperations(nums []int, k int) int {
	var xor = 0
	for _, num := range nums {
		xor ^= num
	}
	if xor == k {
		return 0
	}
	xor ^= k
	var flips = 0
	for xor > 0 {
		if xor&1 == 1 {
			flips++
		}
		xor >>= 1
	}
	return flips
}

// func minOperations(nums []int, k int) int {
// 	for _, num := range nums {
// 		k ^= num
// 	}
// 	return bits.OnesCount(uint(k))
// }
