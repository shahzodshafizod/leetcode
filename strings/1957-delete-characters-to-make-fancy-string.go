package strings

// https://leetcode.com/problems/delete-characters-to-make-fancy-string/

// Approach: Stack
// Time: O(n)
// Space: O(n)
func makeFancyString(s string) string {
	var stack = make([]rune, 0, len(s))
	var size = 0
	for _, c := range s {
		if size < 2 || stack[size-1] != c || stack[size-2] != c {
			stack = append(stack, c)
			size++
		}
	}
	return string(stack)
}

// // Approach: Two Pointers
// // Time: O(n)
// // Space: O(n)
// func makeFancyString(s string) string {
// 	var n = len(s)
// 	var fancy = []byte(s)
// 	var slow = 2
// 	for fast := 2; fast < n; fast++ {
// 		if fancy[fast] != fancy[slow-1] || fancy[fast] != fancy[slow-2] {
// 			fancy[slow] = s[fast]
// 			slow++
// 		}
// 	}
// 	return string(fancy[:min(slow, len(s))])
// }
