package strings

// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

func finalValueAfterOperations(operations []string) int {
	var res int

	for _, op := range operations {
		if op[1] == '+' {
			res++
		} else {
			res--
		}
	}

	return res
}
