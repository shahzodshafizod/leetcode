package dp

import "strings"

// https://leetcode.com/problems/stickers-to-spell-word/

// Approach: Top-Down Dynamic Programming
// N=len(stickers)
// M=len(target)
// for every letter in target we decide whether to choose the sticker or not
// in the worst case for every letter in target we choose every sticker
// Time: O(2^N ∗M)
// Space: O(2^N ∗M)
func minStickers(stickers []string, target string) int {
	const MAX_INT = 16
	dict := make([][26]int, len(stickers))
	for idx, sticker := range stickers {
		for _, c := range sticker {
			dict[idx][c-'a']++
		}
	}
	memo := make(map[string]int)
	var dfs func(target string) int
	dfs = func(target string) int {
		if target == "" {
			return 0
		}
		if memo[target] != 0 {
			return memo[target]
		}
		var targetMap [26]int
		for _, c := range target {
			targetMap[c-'a']++
		}
		count := MAX_INT
		for _, sticker := range dict {
			if sticker[target[0]-'a'] == 0 {
				continue
			}
			next := ""
			for idx, cnt := range targetMap {
				if cnt > sticker[idx] {
					next += strings.Repeat(string(byte('a'+idx)), cnt-sticker[idx])
				}
			}
			count = min(count, 1+dfs(next)) // +1 mean we chose this sticker
		}
		memo[target] = count
		return count
	}
	count := dfs(target)
	if count == MAX_INT {
		count = -1
	}
	return count
}
