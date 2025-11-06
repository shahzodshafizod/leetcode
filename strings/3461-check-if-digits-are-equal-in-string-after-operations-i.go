package strings

// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/

func hasSameDigits(s string) bool {
	n := len(s)

	digit := make([]int, n)
	for i, c := range s {
		digit[i] = int(c) - int('0')
	}

	for size := n - 1; size > 1; size-- {
		for i := range size {
			digit[i] = (digit[i] + digit[i+1]) % 10
		}
	}

	return digit[0] == digit[1]
}
