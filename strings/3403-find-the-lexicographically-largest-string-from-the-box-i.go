package strings

// https://leetcode.com/problems/find-the-lexicographically-largest-string-from-the-box-i/

// Approach 1: Enumeration
// Time: O(nk), n=len(word), k=n-numFriends+1
// Space: O(1)
func answerString(word string, numFriends int) string {
	if numFriends <= 1 {
		return word
	}
	var n = len(word)
	var k = n - numFriends + 1
	var largest = ""
	for i := 0; i < n; i++ {
		largest = max(largest, word[i:min(i+k, n)])
	}
	return largest
}
