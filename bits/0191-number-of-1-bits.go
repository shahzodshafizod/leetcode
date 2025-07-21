package bits

// https://leetcode.com/problems/number-of-1-bits/

// Brian Kernighan's algorithm
// turning off its rightmost set bit
// 0b00000000000000000000000000001011
// 0b00000000000000000000000000001010
// 0b00000000000000000000000000001000
// 0b00000000000000000000000000000000
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		num &= num - 1
		count++
	}
	return count
}

// func hammingWeight(num uint32) int {
// 	return bits.OnesCount(uint(num)) // math/bits
// }
