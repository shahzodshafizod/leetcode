package strings

import "strconv"

// https://leetcode.com/problems/fizz-buzz/

// Approach #3. the use of a modulus operator may have "some" impact on time complexity
// Time: O(N)
// Space: O(1)
func fizzBuzz(n int) []string {
	answer := make([]string, n)

	fizz, buzz := 3, 5
	for idx := 1; idx <= n; idx++ {
		if idx == fizz {
			answer[idx-1] = "Fizz"
			fizz += 3
		}

		if idx == buzz {
			answer[idx-1] += "Buzz"
			buzz += 5
		}

		if len(answer[idx-1]) == 0 {
			answer[idx-1] = strconv.Itoa(idx)
		}
	}

	return answer
}

// // Approach #2
// // Time: O(N)
// // Space: O(1)
// func fizzBuzz(n int) []string {
// 	var source = [4]string{"", "Fizz", "Buzz", "FizzBuzz"}
// 	var answer = make([]string, n)
// 	var mask int
// 	for idx := 1; idx <= n; idx++ {
// 		mask = 0
// 		if idx%3 == 0 {
// 			mask++ // +01
// 		}
// 		if idx%5 == 0 {
// 			mask += 2 // +10
// 		}
// 		if mask == 0 {
// 			answer[idx-1] = strconv.Itoa(idx)
// 		} else {
// 			answer[idx-1] = source[mask]
// 		}
// 	}
// 	return answer
// }

// // Approach #1
// // Time: O(N)
// // Space: O(1)
// func fizzBuzz(n int) []string {
// 	var answer = make([]string, n)
// 	for idx := 1; idx <= n; idx++ {
// 		if idx%3 == 0 {
// 			answer[idx-1] = "Fizz"
// 		}
// 		if idx%5 == 0 {
// 			answer[idx-1] += "Buzz"
// 		}
// 		if answer[idx-1] == "" {
// 			answer[idx-1] = strconv.Itoa(idx)
// 		}
// 	}
// 	return answer
// }
