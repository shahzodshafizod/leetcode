package strings

// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/

func getLucky(s string, k int) int {
	var number = 0
	var temp int
	for _, c := range s {
		temp = int(c-'a') + 1
		number += temp/10 + temp%10
	}
	for k--; k > 0; k-- {
		temp, number = number, 0
		for temp > 0 {
			number += temp % 10
			temp /= 10
		}
	}
	return number
}
