package stacks

// https://leetcode.com/problems/make-the-string-great/

func makeGood(s string) string {
	var st = make([]byte, 0)
	for idx := range s {
		if len(st) > 0 && (st[len(st)-1]-s[idx] == 32 || s[idx]-st[len(st)-1] == 32) {
			st = st[:len(st)-1]
		} else {
			st = append(st, s[idx])
		}
	}
	return string(st)
}

// func makeGood(s string) string {
// 	var diff = func(x uint8, y uint8) uint8 {
// 		if x < y {
// 			x, y = y, x
// 		}
// 		return x - y
// 	}
// 	var st = make([]byte, 0)
// 	for idx := range s {
// 		if len(st) > 0 && diff(st[len(st)-1], s[idx]) == 32 {
// 			st = st[:len(st)-1]
// 		} else {
// 			st = append(st, s[idx])
// 		}
// 	}
// 	return string(st)
// }
