package greedy

// https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/

// Time: O(log10(num))
// Space: O(1)
func maxDiff(num int) int {
	var digit int

	zero, non0and1, nine := 0, 0, 9

	for tmp := num; tmp > 0; tmp /= 10 {
		digit = tmp % 10
		zero = digit

		if digit != 0 && digit != 1 {
			non0and1 = digit
		}

		if digit != 9 {
			nine = digit
		}
	}

	var zeros, nines [10]int
	for digit = range 10 {
		zeros[digit] = digit
		nines[digit] = digit
	}

	if zero != 1 {
		zeros[zero] = 1
	} else {
		zeros[non0and1] = 0
	}

	nines[nine] = 9
	maximum, minimum := 0, 0
	decimal := 1

	for tmp := num; tmp > 0; tmp /= 10 {
		digit = tmp % 10
		maximum += nines[digit] * decimal
		minimum += zeros[digit] * decimal
		decimal *= 10
	}

	return maximum - minimum
}
