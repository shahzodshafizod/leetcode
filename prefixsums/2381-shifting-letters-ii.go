package prefixsums

// https://leetcode.com/problems/shifting-letters-ii/

// Approach: Line Sweep
// Time: O(n+m), n=len(s), m=len(shifts)
// Space: O(n)
func shiftingLetters(s string, shifts [][]int) string {
	line := make([]int, len(s)+1)
	var direction int
	for idx := range shifts {
		direction = shifts[idx][2]
		if direction == 0 {
			direction = -1
		}
		line[shifts[idx][0]] += direction
		line[shifts[idx][1]+1] -= direction
	}
	delta := 0
	b := []byte(s)
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
