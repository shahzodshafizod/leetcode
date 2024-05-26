package strings

import "strings"

// https://leetcode.com/problems/student-attendance-record-i/

func checkRecord(s string) bool {
	return strings.Count(s, "A") < 2 &&
		!strings.Contains(s, "LLL")
}

// func checkRecord(s string) bool {
// 	var absences, lates = 0, 0
// 	for _, c := range s {
// 		if c == 'L' {
// 			lates++
// 			if lates == 3 {
// 				return false
// 			}
// 		} else {
// 			if c == 'A' {
// 				absences++
// 				if absences == 2 {
// 					return false
// 				}
// 			}
// 			lates = 0
// 		}
// 	}
// 	return true
// }
