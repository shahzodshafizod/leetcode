package strings

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/count-and-say/

// RLE: Run-length encoding

// Approach: Iterative
// Time: O(nxm), m=max(len(RLE))
// Space: O(n)
func countAndSay(n int) string {
	rle := "1"

	var newRLE strings.Builder
	for ; n > 1; n-- {
		newRLE.Reset()

		count, slen := 1, len(rle)
		for idx := 1; idx < slen; idx++ {
			if rle[idx] != rle[idx-1] {
				newRLE.WriteString(strconv.Itoa(count))
				newRLE.WriteByte(rle[idx-1])

				count = 0
			}

			count++
		}

		newRLE.WriteString(strconv.Itoa(count))
		newRLE.WriteByte(rle[slen-1])
		rle = newRLE.String()
	}

	return rle
}

// // Approach: Recursive
// // Time: O(nxm), m=max(len(RLE))
// // Space: O(n)
// func countAndSay(n int) string {
// 	if n == 1 {
// 		return "1"
// 	}
// 	var s = countAndSay(n - 1)
// 	var rle strings.Builder
// 	var count, slen = 1, len(s)
// 	for idx := 1; idx < slen; idx++ {
// 		if s[idx] != s[idx-1] {
// 			rle.WriteString(strconv.Itoa(count))
// 			rle.WriteByte(s[idx-1])
// 			count = 0
// 		}
// 		count++
// 	}
// 	rle.WriteString(strconv.Itoa(count))
// 	rle.WriteByte(s[slen-1])
// 	return rle.String()
// }
