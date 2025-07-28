package bits

// https://leetcode.com/problems/minimize-xor/

// Approach: Greedy + Bit Manipulation
// Time: O(log n)
// Space: O(1)
func minimizeXor(num1 int, num2 int) int {
	x := num1

	diff := 0
	for num1 > 0 || num2 > 0 {
		diff += (num2 & 1) - (num1 & 1)
		num1 >>= 1
		num2 >>= 1
	}

	bit := 0
	for diff > 0 && bit < 32 {
		if (1<<bit)&x == 0 {
			x ^= 1 << bit
			diff--
		}

		bit++
	}

	for diff < 0 {
		x &= x - 1
		diff++
	}

	return x
}
