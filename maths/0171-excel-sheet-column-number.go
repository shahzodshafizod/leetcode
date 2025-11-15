package maths

// https://leetcode.com/problems/excel-sheet-column-number/

func titleToNumber(columnTitle string) int {
	columnNumber, base := 0, 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		columnNumber += int(columnTitle[i]-'A'+1) * base
		base *= 26
	}

	return columnNumber
}
