package hashes

// https://leetcode.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	var hash = make(map[rune]int)
	var table = make(map[rune]int)
	for _, r := range t {
		table[r]++
	}
	var count = len(table)
	var minLength = 0
	var window string
	var leftIndex = 0
	for index, r := range s {
		if miqdor := table[r]; miqdor > 0 {
			hash[r]++
			if hash[r] == miqdor {
				count--
			}
			if count == 0 {
				rightIndex := index
				for ; leftIndex < rightIndex; leftIndex++ {
					var key = rune(s[leftIndex])
					if count, exists := hash[key]; exists {
						if count == table[key] {
							break
						}
						hash[key]--
					}
				}
				if length := rightIndex - leftIndex + 1; minLength == 0 || length < minLength {
					minLength = length
					window = s[leftIndex : rightIndex+1]
				}
				hash[rune(s[leftIndex])]--
				leftIndex++
				count = 1
			}
		}
	}
	return window
}
