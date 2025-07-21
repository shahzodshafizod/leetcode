package dp

// https://leetcode.com/problems/regular-expression-matching/

func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	cache := make([][]*bool, slen)
	for idx := range cache {
		cache[idx] = make([]*bool, plen)
	}
	var result bool
	var dp func(int, int) bool
	dp = func(sid int, pid int) bool {
		if sid < 0 || pid < 0 {
			for pid > 0 && p[pid] == '*' {
				pid -= 2
			}
			return sid < 0 && pid < 0
		}
		if cache[sid][pid] != nil {
			return *cache[sid][pid]
		}
		if p[pid] == '*' {
			result = (p[pid-1] == '.' || p[pid-1] == s[sid]) &&
				(dp(sid-1, pid) || dp(sid-1, pid-1)) || dp(sid, pid-2)
		} else {
			result = (p[pid] == '.' || p[pid] == s[sid]) && dp(sid-1, pid-1)
		}
		cache[sid][pid] = &result
		return result
	}
	return dp(slen-1, plen-1)
}
