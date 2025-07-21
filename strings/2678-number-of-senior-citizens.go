package strings

// https://leetcode.com/problems/number-of-senior-citizens/

func countSeniors(details []string) int {
	count := 0
	for _, detail := range details {
		if detail[11] > '6' || detail[11] == '6' && detail[12] != '0' {
			count++
		}
		// age, err := strconv.Atoi(detail[11:13])
		// int(detail[11]-'0')*10+int(detail[12]-'0')
	}
	return count
}
