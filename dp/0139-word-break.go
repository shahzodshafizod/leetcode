package dp

// https://leetcode.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	var sn = len(s)
	var dp = make([]bool, sn+1)
	dp[sn] = true
	for idx := sn - 1; idx >= 0; idx-- {
		for _, word := range wordDict {
			if idx+len(word) <= sn && s[idx:idx+len(word)] == word {
				dp[idx] = dp[idx+len(word)]
			}
			if dp[idx] {
				break
			}
		}
	}
	return dp[0]
}

// func wordBreak(s string, wordDict []string) bool {
// 	var wordices = make([][2]int, 0)
// 	var start, next, skips int
// 	var sn = len(s)
// 	var exists = make([]bool, sn)
// 	for _, word := range wordDict {
// 		start = strings.Index(s, word)
// 		skips = 0
// 		for start >= 0 {
// 			start += skips
// 			skips = start + 1
// 			next = start + len(word)
// 			wordices = append(wordices, [2]int{start, next})
// 			for start < next {
// 				exists[start] = true
// 				start++
// 			}
// 			start = strings.Index(s[skips:], word)
// 		}
// 	}
// 	for _, exists := range exists {
// 		if !exists {
// 			return false
// 		}
// 	}
// 	var dfs func(idx int) bool
// 	dfs = func(idx int) bool {
// 		if idx >= sn {
// 			return idx == sn
// 		}
// 		for _, index := range wordices {
// 			if index[0] == idx {
// 				if dfs(index[1]) {
// 					return true
// 				}
// 			}
// 		}
// 		return false
// 	}
// 	return dfs(0)
// }
