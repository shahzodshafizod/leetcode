package strings

// https://leetcode.com/problems/integer-to-english-words/

var dictionary = map[int]string{
	0: "Zero", 1: "One", 2: "Two", 3: "Three", 4: "Four",
	5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine",
	10: "Ten", 11: "Eleven", 12: "Twelve", 13: "Thirteen", 14: "Fourteen",
	15: "Fifteen", 16: "Sixteen", 17: "Seventeen", 18: "Eighteen", 19: "Nineteen",
	20: "Twenty", 30: "Thirty", 40: "Forty", 50: "Fifty", 60: "Sixty",
	70: "Seventy", 80: "Eighty", 90: "Ninety", 100: "Hundred", 1_000: "Thousand",
	1_000_000: "Million", 1_000_000_000: "Billion", 1_000_000_000_000: "Trillion",
}

var groups = []int{
	1_000_000_000_000,
	1_000_000_000,
	1_000_000,
	1_000,
	100,
	20,
}

func numberToWords(num int) string {
	if num <= 20 {
		return dictionary[num]
	}
	for _, group := range groups {
		if num >= group {
			var words string

			if group == 20 {
				group = 10
				words = dictionary[num/group*10]
			} else {
				words = numberToWords(num/group) + " " + dictionary[group]
			}

			if num%group != 0 {
				words += " " + numberToWords(num%group)
			}

			return words
		}
	}
	return ""
}

// Follow-up: Convert NUMBER in reverse order to WORDS without reversing the NUMBER itself

func numberToWordsReverse(num int) string {
	var divider = 1
	var tempNum = num
	for tempNum > 0 {
		tempNum /= 10
		divider *= 10
	}
	return numberToWordsReverseHelper(num, divider, 0)
}

func numberToWordsReverseHelper(num int, divider int, level int) string {
	var words string
	var current = divider + num
	divider /= 1000
	if divider != 0 {
		current /= divider
		if num%divider != 0 {
			words += numberToWordsReverseHelper(num%divider, divider, level+1) + " "
		}
	}
	num = 0           // reverse every block separately
	for current > 1 { // skip the first digit and reverse: 1004 => 400, 145 => 54
		num = num*10 + current%10
		current /= 10
	}
	words += numberToWords(num)
	if levels[level] != "" {
		words += " " + levels[level]
	}
	return words
}

var levels = []string{"", "Thousand", "Million", "Billion", "Trillion"}
