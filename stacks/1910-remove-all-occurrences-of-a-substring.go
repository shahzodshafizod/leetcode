package stacks

// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/

// Approach: KMP (Knuth-Morris-Pratt) Algorithm
// Time: O(n+m), n=len(s), m=len(part)
// Space: O(n+m)
func removeOccurrences(s string, part string) string {
	var pn = len(part)
	var getLPS = func(part string) []int {
		var lps = make([]int, pn)
		for preflen, idx := 0, 1; idx < pn; idx++ {
			if part[idx] == part[preflen] {
				preflen++
				lps[idx] = preflen
			} else if preflen > 0 {
				preflen = lps[preflen-1]
				idx--
			}
		}
		return lps
	}
	var lps = getLPS(part)
	var pid, sid, sn = 0, 0, len(s)
	var stack = make([]byte, sn)
	var nextID = make([]int, sn)
	var diff = 0
	for sid < sn {
		stack[sid-diff] = s[sid]
		if s[sid] == part[pid] {
			pid++
			nextID[sid-diff] = pid
			if pid == pn {
				diff += pn
				pid = 0
				if sid >= diff {
					pid = nextID[sid-diff]
				}
			}
		} else if pid > 0 {
			pid = lps[pid-1]
			sid--
		} else {
			nextID[sid-diff] = 0
		}
		sid++
	}
	return string(stack[:sn-diff])
}
