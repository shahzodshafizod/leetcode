package maths

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/fraction-to-recurring-decimal/

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var sb strings.Builder
	if (numerator < 0) != (denominator < 0) {
		_ = sb.WriteByte('-')
	}

	toInt64Abs := func(x int) int64 {
		if x < 0 {
			return -int64(x)
		}

		return int64(x)
	}
	n, d := toInt64Abs(numerator), toInt64Abs(denominator)
	_, _ = sb.WriteString(strconv.FormatInt(n/d, 10))

	remainder := n % d
	if remainder == 0 {
		return sb.String()
	}

	_ = sb.WriteByte('.')

	fractions := map[int64]int{}
	for remainder != 0 {
		if p, found := fractions[remainder]; found {
			s := sb.String()

			return s[:p] + "(" + s[p:] + ")"
		}

		fractions[remainder] = sb.Len()
		remainder *= 10
		_, _ = sb.WriteString(strconv.FormatInt(remainder/d, 10))
		remainder %= d
	}

	return sb.String()
}
