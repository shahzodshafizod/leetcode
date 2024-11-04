package strings

import "strings"

// https://leetcode.com/problems/circular-sentence/

// Time: O(n)
// Space: O(n)
func isCircularSentence(sentence string) bool {
	var words = strings.Split(sentence, " ")
	var n = len(words)
	for idx := 1; idx < n; idx++ {
		if words[idx][0] != words[idx-1][len(words[idx-1])-1] {
			return false
		}
	}
	return words[0][0] == words[n-1][len(words[n-1])-1]
}

// // Time: O(n)
// // Space: O(1)
// func isCircularSentence(sentence string) bool {
// 	for i := range sentence {
// 		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
// 			return false
// 		}
// 	}
// 	return sentence[0] == sentence[len(sentence)-1]
// }
