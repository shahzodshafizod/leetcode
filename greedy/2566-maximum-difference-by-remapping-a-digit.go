package greedy

// https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/

// Time: O(log10(num))
// Space: O(1)
func minMaxDifference(num int) int {
	var zeros, nines [10]int
	var digit int
	for digit = 0; digit < 10; digit++ {
		zeros[digit] = digit
		nines[digit] = digit
	}
	var tmp = num
	var zero, nine = 0, 9
	for ; tmp > 0; tmp /= 10 {
		digit = tmp % 10
		zero = digit
		if digit != 9 {
			nine = digit
		}
	}
	zeros[zero] = 0
	nines[nine] = 9
	var max, min = 0, 0
	tmp = num
	var decimal = 1
	for ; tmp > 0; tmp /= 10 {
		digit = tmp % 10
		max += nines[digit] * decimal
		min += zeros[digit] * decimal
		decimal *= 10
	}
	return max - min
}
