package hashes

// https://leetcode.com/problems/roman-to-integer/

func romanToInt(s string) int {
	var table = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var num = 0
	var next byte = 'I'
	var curr byte
	for idx := len(s) - 1; idx >= 0; idx-- {
		curr = s[idx]
		if table[curr] < table[next] {
			num -= table[curr]
		} else {
			num += table[curr]
		}
		next = curr
	}
	return num
}
