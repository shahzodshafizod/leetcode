package bits

// https://leetcode.com/problems/maximum-odd-binary-number/

// Greedy Bit Manipulation
func maximumOddBinaryNumber(s string) string {
	const ZERO, ONE = '0', '1'
	n := len(s)
	b := []byte(s)
	left, right := 0, n-1
	for left <= right {
		switch {
		case b[left] == ONE:
			left++
		case b[right] == ZERO:
			right--
		default:
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
	}
	left--
	b[left], b[n-1] = b[n-1], b[left]
	return string(b)
}
