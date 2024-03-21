package dp

// https://leetcode.com/problems/regular-expression-matching/

// // mine: (3ms) vs NeetCode (222ms)
func isMatch(s string, p string) bool {
	var slen, plen = len(s), len(p)
	var cache = make([][]*bool, slen)
	for idx := range cache {
		cache[idx] = make([]*bool, plen)
	}
	var dp func(int, int) bool
	dp = func(sid int, pid int) bool {
		if sid < 0 && pid < 0 {
			return true
		}
		if sid < 0 {
			if p[pid] == '*' {
				return dp(sid, pid-2)
			}
			return false
		}
		if pid < 0 {
			return false
		}
		if cache[sid][pid] != nil {
			return *cache[sid][pid]
		}
		var result bool
		switch p[pid] {
		case '.':
			result = dp(sid-1, pid-1)
		case '*':
			var match = true
			if sid >= 0 {
				match = p[pid-1] == '.' || p[pid-1] == s[sid]
			}
			result = match && (dp(sid-1, pid) || dp(sid-1, pid-1)) || dp(sid, pid-2)
		default:
			result = s[sid] == p[pid] && dp(sid-1, pid-1)
		}
		cache[sid][pid] = &result
		return result
	}
	return dp(slen-1, plen-1)
}

// // NeetCode (222ms)
// func isMatch(s string, p string) bool {
// 	var slen, plen = len(s), len(p)
// 	var dp func(int, int) bool
// 	dp = func(sid int, pid int) bool {
// 		if sid >= slen && pid >= plen {
// 			return true
// 		}
// 		if pid >= plen {
// 			return false
// 		}
// 		match := sid < slen && (s[sid] == p[pid] || p[pid] == '.')
// 		if (pid+1) < plen && p[pid+1] == '*' {
// 			return dp(sid, pid+2) || match && dp(sid+1, pid)
// 		}
// 		if match {
// 			return dp(sid+1, pid+1)
// 		}
// 		return false
// 	}
// 	return dp(0, 0)
// }
