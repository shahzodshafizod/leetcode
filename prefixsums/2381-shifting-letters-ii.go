package prefixsums

// https://leetcode.com/problems/shifting-letters-ii/

// Approach: Line Sweep
// Time: O(n+m), n=len(s), m=len(shifts)
// Space: O(m)
func shiftingLetters(s string, shifts [][]int) string {
	var line = make([]int, len(s)+1)
	var start, end, direction int
	for idx := range shifts {
		start = shifts[idx][0]
		end = shifts[idx][1]
		direction = shifts[idx][2]
		if direction == 0 {
			direction = -1
		}
		line[start] += direction
		line[end+1] -= direction
	}
	var delta = 0
	var b = []byte(s)
	var code int
	for idx := range b {
		delta += line[idx]
		code = (int(b[idx]-'a') + delta) % 26
		if code < 0 {
			code += 26
		}
		b[idx] = byte('a' + code)
	}
	return string(b)
}
