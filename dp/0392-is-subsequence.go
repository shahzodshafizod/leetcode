package dp

/*
Follow up:
Suppose there are lots of incoming s, say s1, s2, ..., sk
where k >= 109, and you want to check one by one to see if t has its subsequence.
In this scenario, how would you change your code?

Answer:
Create a trie of all s's with the following node struct:
type Node struct {
	char byte
	count int
}
When creating the trie of s's, every time you pass through a node,
increment its count, it means how many s's pass through this node.
Then iterate over t and return the count of the last station if you find it.
*/

// https://leetcode.com/problems/is-subsequence/

// time: O(len(t))
// space: O(1)
func isSubsequence(s string, t string) bool {
	sn, tn := len(s), len(t)
	si := 0
	for ti := 0; ti < tn && si < sn; ti++ {
		if s[si] == t[ti] {
			si++
		}
	}
	return si == sn
}

// // time: O(len(t))
// // space: O(len(t)) for recursion stack
// func isSubsequence(s string, t string) bool {
// 	var tlen, slen = len(t), len(s)
// 	var dfs func(si int, ti int) bool
// 	dfs = func(si int, ti int) bool {
// 		if ti == tlen || si == slen {
// 			return si == slen
// 		}
// 		// decision to include s[i]
// 		if s[si] == t[ti] {
// 			return dfs(si+1, ti+1)
// 		}
// 		// decision NOT to include s[i]
// 		return dfs(si, ti+1)
// 	}
// 	return dfs(0, 0)
// }
