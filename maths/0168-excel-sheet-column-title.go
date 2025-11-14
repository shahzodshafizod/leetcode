package maths

import (
	"slices"
	"strings"
)

// https://leetcode.com/problems/excel-sheet-column-title/

func convertToTitle(columnNumber int) string {
	var sb strings.Builder

	for columnNumber > 0 {
		columnNumber--
		_ = sb.WriteByte(byte(int('A') + columnNumber%26))
		columnNumber /= 26
	}

	columnTitle := []byte(sb.String())
	slices.Reverse(columnTitle)

	return string(columnTitle)
}
