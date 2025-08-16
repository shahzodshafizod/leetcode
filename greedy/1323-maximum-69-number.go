package greedy

// https://leetcode.com/problems/maximum-69-number/

func maximum69Number(num int) int {
	dec, pos := 1, 0

	for tmp := num; tmp > 0; tmp /= 10 {
		if tmp%10 == 6 {
			pos = dec
		}

		dec *= 10
	}

	return num + 3*pos
}
