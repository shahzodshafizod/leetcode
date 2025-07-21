package bits

// https://leetcode.com/problems/single-number-iii/

// time: O(n)
// space: O(1)
func singleNumberIII(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	bit := 1
	for bit&xor == 0 {
		bit <<= 1
	}
	a := 0
	for _, num := range nums {
		if num&bit == bit {
			a ^= num
		}
	}
	b := xor ^ a
	return []int{a, b}
}
