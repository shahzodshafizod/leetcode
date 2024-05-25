package backtracking

import "strings"

// https://leetcode.com/problems/word-break-ii/

func wordBreak(s string, wordDict []string) []string {
	type index struct {
		start int
		next  int
		index int
	}
	var wordices = make([]*index, 0)
	var start, next, skips int
	var sn = len(s)
	var exists = make([]bool, sn)
	for idx, word := range wordDict {
		start = strings.Index(s, word)
		skips = 0
		for start >= 0 {
			start += skips
			skips = start + 1
			next = start + len(word)
			wordices = append(wordices, &index{start: start, next: next, index: idx})
			for start < next {
				exists[start] = true
				start++
			}
			start = strings.Index(s[skips:], word)
		}
	}
	for _, exists := range exists {
		if !exists {
			return []string{}
		}
	}
	var sentences = make([]string, 0)
	var dfs func(idx int, sentence *string)
	dfs = func(idx int, sentence *string) {
		if idx == sn {
			sentences = append(sentences, (*sentence)[1:])
			return
		}
		for _, index := range wordices {
			if index.start == idx {
				*sentence += " " + wordDict[index.index]
				dfs(index.next, sentence)
				*sentence = (*sentence)[:strings.LastIndex(*sentence, wordDict[index.index])-1]
			}
		}
	}
	var sentence = ""
	dfs(0, &sentence)
	return sentences
}

// // DP: Tabulation
// func wordBreak(s string, wordDict []string) []string {
// 	var sn = len(s)
// 	var sentences = make([][]string, sn+1)
// 	sentences[sn] = []string{""}
// 	sentences[0] = []string{}
// 	for idx := sn - 1; idx >= 0; idx-- {
// 		for _, word := range wordDict {
// 			if idx+len(word) <= sn && s[idx:idx+len(word)] == word {
// 				for _, tail := range sentences[idx+len(word)] {
// 					if tail != "" {
// 						tail = " " + tail
// 					}
// 					sentences[idx] = append(sentences[idx], word+tail)
// 				}
// 			}
// 		}
// 	}
// 	return sentences[0]
// }
