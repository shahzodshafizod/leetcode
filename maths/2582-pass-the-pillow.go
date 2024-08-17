package maths

// https://leetcode.com/problems/pass-the-pillow/

// Approach #2: Math
func passThePillow(n int, time int) int {
	rounds := time / (n - 1)
	extras := time % (n - 1)
	if rounds&1 == 0 {
		return extras + 1
	}
	return n - extras
}

// // Approach #1: Simulation
// func passThePillow(n int, time int) int {
// 	var person, direction = 1, 1
// 	for time > 0 {
// 		time--
// 		person += direction
// 		if person == 1 || person == n {
// 			direction *= -1
// 		}
// 	}
// 	return person
// }
