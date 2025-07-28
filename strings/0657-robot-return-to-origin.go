package strings

// https://leetcode.com/problems/robot-return-to-origin/

// Approach: Simulation
// Time: O(n)
// Space: O(1)
func judgeCircle(moves string) bool {
	x, y := 0, 0

	for _, move := range moves {
		switch move {
		case 'R':
			y++
		case 'L':
			y--
		case 'D':
			x++
		case 'U':
			x--
		}
	}

	return x == 0 && y == 0
}

// // Approach: Counting
// func judgeCircle(moves string) bool {
// 	var count = map[rune]int{'R': 0, 'L': 0, 'U': 0, 'D': 0}
// 	for _, move := range moves {
// 		count[move]++
// 	}
// 	return count['R'] == count['L'] && count['U'] == count['D']
// }
