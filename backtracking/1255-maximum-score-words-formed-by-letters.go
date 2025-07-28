package backtracking

// https://leetcode.com/problems/maximum-score-words-formed-by-letters/

func maxScoreWords(words []string, letters []byte, score []int) int {
	var lmap [26]int
	for _, letter := range letters {
		lmap[letter-'a']++
	}

	getScore := func(idx int) int {
		var c int

		var count [26]int

		wordScore := 0

		for _, letter := range words[idx] {
			c = int(letter - 'a')
			wordScore += score[c]

			count[c]++
			if count[c] > lmap[c] {
				return -1
			}
		}

		return wordScore
	}

	var dfs func(idx int) int

	dfs = func(idx int) int {
		if idx < 0 {
			return 0
		}
		// decision NOT to include words[idx]
		maxScore := dfs(idx - 1)
		// decision to include words[idx]
		if wordScore := getScore(idx); wordScore != -1 {
			for _, letter := range words[idx] {
				lmap[letter-'a']--
			}

			maxScore = max(maxScore, wordScore+dfs(idx-1))
			// backtracking...
			for _, letter := range words[idx] {
				lmap[letter-'a']++
			}
		}

		return maxScore
	}

	return dfs(len(words) - 1)
}
