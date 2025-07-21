package bits

/*
Note:
This question is the same as 1009:
https://leetcode.com/problems/complement-of-base-10-integer/
*/

// https://leetcode.com/problems/number-complement/

func findComplement(num int) int {
	bit := 1
	for bit <= num {
		num ^= bit
		bit <<= 1
	}
	return num
}
